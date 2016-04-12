package storage

import (
	"os"

	"github.com/jmoiron/sqlx"
	// Load Postgres
	_ "github.com/lib/pq"
)

var connectionDSN = os.Getenv("MUNI_POSTGRES_DSN")

// CreateConnection connects to a Postgres DB using the DSN provided
func CreateConnection() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", connectionDSN)

	if err != nil {
		return nil, err
	}

	return db, nil
}
