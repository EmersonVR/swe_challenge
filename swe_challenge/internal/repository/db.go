package repository

import (
    "database/sql"
    _ "github.com/lib/pq" // Import para registrar el driver Postgres
    "fmt"
)

// NewDBConnection crea una conexión a la base de datos CockroachDB
func NewDBConnection(host string, port int, user, password, dbName string) (*sql.DB, error) {
    // sslmode puede variar: 'disable', 'require', 'verify-full', etc. 
    // Ajusta según la configuración de tu cluster en CockroachDB.
    connStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=require",
        user, password, host, port, dbName,
    )

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    // Verificar la conexión realizando un ping
    if err := db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}
