package test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"swe_challenge/internal/repository"
)

func TestGetAllAcciones(t *testing.T) {
	// Simulación de base de datos con sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creando sqlmock: %v", err)
	}
	defer db.Close()

	// Simulación de datos devueltos por la consulta
	rows := sqlmock.NewRows([]string{
		"id", "ticker", "objetivo_desde", "objetivo_a", "empresa", 
		"accion", "corretaje", "rating_from", "rating_to", "hora",
	}).AddRow(1, "AAPL", 150.00, 160.00, "Apple Inc.", "Buy", "Goldman Sachs", "Buy", "Strong Buy", time.Now())

	// Simular la consulta que usa GetAll
	mock.ExpectQuery(`SELECT id, ticker, objetivo_desde, objetivo_a, empresa, accion, corretaje, rating_from, rating_to, hora 
                      FROM acciones 
                      ORDER BY hora DESC 
                      LIMIT \$1 OFFSET \$2`).
		WillReturnRows(rows)

	// Crear repositorio con la BD simulada
	repo := repository.NewAccionRepository(db)

	// Ejecutar la función
	result, err := repo.GetAll(10, 0)
	if err != nil {
		t.Fatalf("Error en GetAll: %v", err)
	}

	// Verificar que la consulta devolvió datos
	if len(result) != 1 {
		t.Fatalf("Se esperaba 1 acción, pero se obtuvieron %d", len(result))
	}

	// Verificar que los datos son correctos
	expectedTicker := "AAPL"
	if result[0].Ticker != expectedTicker {
		t.Errorf("Se esperaba ticker %s, pero se obtuvo %s", expectedTicker, result[0].Ticker)
	}

	// Verificar que todas las expectativas de la BD se cumplieron
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("No se cumplieron todas las expectativas de la BD: %v", err)
	}
}
