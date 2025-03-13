package test

import (
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestDBConnection(t *testing.T) {
	// Simulamos la base de datos con sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creando sqlmock: %v", err)
	}
	defer db.Close()

	// Simulamos la consulta "SELECT 1"
	mock.ExpectQuery("SELECT 1").WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))

	// Ahora intentamos ejecutar esa consulta
	var result int
	err = db.QueryRow("SELECT 1").Scan(&result)
	if err != nil {
		t.Fatalf("Error ejecutando SELECT 1: %v", err)
	}

	// Validamos que el resultado sea correcto
	if result != 1 {
		t.Errorf("Se esperaba 1, pero se obtuvo %d", result)
	}

	// Verificamos que todas las expectativas de sqlmock se cumplieron
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("No se cumplieron todas las expectativas de la BD: %v", err)
	}
}
