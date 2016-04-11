package storage

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var connectionDSN string = os.Getenv("MUNI_POSTGRES_DSN")

func CreateConnection() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", connectionDSN)

	if err != nil {
		return nil, err
	}

	return db, nil
}
