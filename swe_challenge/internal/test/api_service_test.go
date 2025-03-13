package test

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"swe_challenge/internal/api"
	"swe_challenge/internal/repository"
	"github.com/DATA-DOG/go-sqlmock"
	"time"
)

// Accion estructura para verificar la respuesta JSON
type Accion struct {
	Ticker        string  `json:"ticker"`
	ObjetivoDesde float64 `json:"objetivo_desde"`
	ObjetivoA     float64 `json:"objetivo_a"`
	Empresa       string  `json:"empresa"`
	Accion        string  `json:"accion"`
	Corretaje     string  `json:"corretaje"`
	RatingFrom    string  `json:"rating_from"`
	RatingTo      string  `json:"rating_to"`
	Hora          time.Time `json:"hora"`
}

func TestGetAccionesEndpoint(t *testing.T) {
	// Simulamos la BD con sqlmock
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Simulamos una respuesta de la base de datos con datos válidos
	rows := sqlmock.NewRows([]string{"id", "ticker", "objetivo_desde", "objetivo_a", "empresa", "accion", "corretaje", "rating_from", "rating_to", "hora"}).
		AddRow(1, "AAPL", 150.00, 160.00, "Apple Inc.", "Buy", "Goldman Sachs", "Buy", "Strong Buy", time.Now())

	mock.ExpectQuery(`SELECT id, ticker, objetivo_desde, objetivo_a, empresa, accion, corretaje, rating_from, rating_to, hora FROM acciones`).
		WillReturnRows(rows)

	// Creamos el repositorio con la BD simulada
	accionRepo := repository.NewAccionRepository(db)

	// Creamos el router con el repositorio
	router := api.NewRouter(accionRepo)

	// Simulamos una petición GET a /acciones
	req, _ := http.NewRequest("GET", "/acciones", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Esperamos una respuesta 200
	if rr.Code != http.StatusOK {
		t.Errorf("Código de estado incorrecto: esperado %d, pero se obtuvo %d", http.StatusOK, rr.Code)
	}

	// Verificamos si la respuesta JSON es válida
	var response []repository.Accion
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error al parsear JSON de respuesta: %v", err)
	}

	// Verifica que se cumplieron todas las expectativas del mock
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("No se cumplieron todas las expectativas de la BD: %v", err)
	}
}
