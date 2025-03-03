package database

import (
	"database/sql"
	"errors"

	"github.com/guilherme-or/go-cleanarch-starter/internal/infra/config"
	_ "github.com/lib/pq"
)

func PostgreSQLConn() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.E.Database.DSN)
	if err != nil {
		return nil, errors.New("database connection open err: " + err.Error())
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.New("database connection ping err: " + err.Error())
	}

	return db, nil
}
