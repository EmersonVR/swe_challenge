package repository

import (
	"database/sql"
	"strconv"
	"strings"
	"time"
	"sort"
	"swe_challenge/internal/services"
)

// AccionRepository encapsula la conexión a la base de datos y las operaciones relacionadas con la tabla 'acciones'.
type AccionRepository struct {
	db *sql.DB
}

// NewAccionRepository crea una nueva instancia de AccionRepository usando la conexión proporcionada.
func NewAccionRepository(db *sql.DB) *AccionRepository {
	return &AccionRepository{db: db}
}

// AccionExists verifica si ya existe un registro en la tabla 'acciones' con el mismo ticker y la misma hora (formato RFC3339).
func (r *AccionRepository) AccionExists(item services.Item) (bool, error) {
	parsedTime, err := time.Parse(time.RFC3339, item.Hora)
	if err != nil {
		return false, err
	}
	var count int
	err = r.db.QueryRow("SELECT COUNT(*) FROM acciones WHERE ticker = $1 AND hora = $2", item.Ticker, parsedTime).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// cleanNumber elimina comas y signos de dólar de un número antes de convertirlo a float64.
func cleanNumber(s string) (float64, error) {
	s = strings.ReplaceAll(s, ",", "") // Eliminar comas
	return strconv.ParseFloat(strings.TrimPrefix(s, "$"), 64)
}

// InsertAccion inserta un nuevo registro en la tabla 'acciones' si no existe ya uno con el mismo ticker y hora.
func (r *AccionRepository) InsertAccion(item services.Item) (bool, error) {
	// Verificar si ya existe un registro con el mismo ticker y hora
	exists, err := r.AccionExists(item)
	if err != nil {
		return false, err
	}
	if exists {
		return false, nil // Los datos ya existen, se omite la inserción
	}

	// Convertir valores de objetivo a float64
	objetivoDesde, err := cleanNumber(item.ObjetivoDesde)
	if err != nil {
		return false, err
	}

	objetivoA, err := cleanNumber(item.ObjetivoA)
	if err != nil {
		return false, err
	}

	// Parsear la hora del item usando el formato RFC3339
	parsedTime, err := time.Parse(time.RFC3339, item.Hora)
	if err != nil {
		return false, err
	}

	// Consulta SQL para insertar el registro en la tabla 'acciones'
	query := `INSERT INTO acciones (
		ticker, objetivo_desde, objetivo_a, empresa, accion, corretaje,
		rating_from, rating_to, hora
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err = r.db.Exec(query,
		item.Ticker,
		objetivoDesde,
		objetivoA,
		item.Empresa,
		item.Accion,
		item.Corretaje,
		item.RatingFrom,
		item.RatingTo,
		parsedTime,
	)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Accion representa el registro de la tabla 'acciones'.
type Accion struct {
	ID            int       `json:"id"`
	Ticker        string    `json:"ticker"`
	ObjetivoDesde float64   `json:"objetivo_desde"`
	ObjetivoA     float64   `json:"objetivo_a"`
	Empresa       string    `json:"empresa"`
	Accion        string    `json:"accion"`
	Corretaje     string    `json:"corretaje"`
	RatingFrom    string    `json:"rating_from"`
	RatingTo      string    `json:"rating_to"`
	Hora          time.Time `json:"hora"`
}

// GetAll consulta y retorna todos los registros de la tabla 'acciones'.
func (r *AccionRepository) GetAll(limit, offset int) ([]Accion, error) {
	rows, err := r.db.Query(`
        SELECT id, ticker, objetivo_desde, objetivo_a, empresa, accion, corretaje, rating_from, rating_to, hora 
        FROM acciones 
        ORDER BY hora DESC
        LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var acciones []Accion
	for rows.Next() {
		var a Accion
		err = rows.Scan(&a.ID, &a.Ticker, &a.ObjetivoDesde, &a.ObjetivoA, &a.Empresa, &a.Accion, &a.Corretaje, &a.RatingFrom, &a.RatingTo, &a.Hora)
		if err != nil {
			return nil, err
		}
		acciones = append(acciones, a)
	}
	return acciones, nil
}

// Acción con puntuación para recomendación
type AccionRecomendada struct {
	Ticker        string    `json:"ticker"`
	Empresa       string    `json:"empresa"`
	Corretaje     string    `json:"corretaje"`
	Accion        string    `json:"accion"`
	RatingFrom    string    `json:"rating_from"`
	RatingTo      string    `json:"rating_to"`
	ObjetivoDesde float64   `json:"objetivo_desde"`
	ObjetivoA     float64   `json:"objetivo_a"`
	Hora          time.Time `json:"hora"`
	Puntuacion    float64   `json:"puntuacion"`
}

// GetRecommendedActions obtiene las mejores acciones según el algoritmo de recomendación
func (r *AccionRepository) GetRecommendedActions() ([]AccionRecomendada, error) {
	rows, err := r.db.Query(`
        SELECT ticker, empresa, corretaje, accion, rating_from, rating_to, 
               objetivo_desde, objetivo_a, hora
        FROM acciones
        WHERE rating_to NOT IN ('Sell', 'Underweight', 'Reduce')
        ORDER BY hora DESC
        LIMIT 50
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var acciones []AccionRecomendada
	for rows.Next() {
		var a AccionRecomendada
		err = rows.Scan(&a.Ticker, &a.Empresa, &a.Corretaje, &a.Accion, &a.RatingFrom,
			&a.RatingTo, &a.ObjetivoDesde, &a.ObjetivoA, &a.Hora)
		if err != nil {
			return nil, err
		}

		// Calculamos el crecimiento esperado
		crecimiento := (a.ObjetivoA - a.ObjetivoDesde) / a.ObjetivoDesde * 100

		// Asignamos un peso a ciertos bancos importantes
		pesoCorretaje := 1.0
		corretajesVIP := map[string]float64{
			"Goldman Sachs": 1.3,
			"Morgan Stanley": 1.2,
			"Citigroup":      1.1,
		}
		if peso, ok := corretajesVIP[a.Corretaje]; ok {
			pesoCorretaje = peso
		}

		// Asignamos un puntaje a la calificación
		puntajeRating := map[string]float64{
			"Buy":            5.0,  
			"Market Perform": 4.0,
			"Neutral":        3.5,
			"Equal Weight":   3.0,
		}
		ratingScore := puntajeRating[a.RatingTo]

		// Cálculo final de la puntuación
		a.Puntuacion = (crecimiento * 2) + (ratingScore * 15) * pesoCorretaje

		acciones = append(acciones, a)
	}

	// Ordenamos de mayor a menor puntuación
	sort.Slice(acciones, func(i, j int) bool {
		return acciones[i].Puntuacion > acciones[j].Puntuacion
	})

	// Retornamos las top 5 acciones
	if len(acciones) >= 20 {
		acciones = acciones[:20]
	}

	return acciones, nil
}
