package database

import (
	"database/sql"

	"github.com/guilherme-or/go-cleanarch-starter/internal/infra/config"
	_ "github.com/lib/pq"
)

func PostgreSQLConn() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.E.Database.DSN)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
