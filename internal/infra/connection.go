package database

import (
	"database/sql"

	"github.com/guilherme-or/go-cleanarch-starter/internal/config"
	_ "github.com/lib/pq"
)

func PostgreSQLConn() *sql.DB {
	db, err := sql.Open("postgres", config.E.Database.DSN)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
