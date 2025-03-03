// Makefile run command entry point
// This file is used to build and run the entire application
package main

import (
	"database/sql"
	"github.com/guilherme-or/go-cleanarch-starter/internal/infra/database"

	"github.com/guilherme-or/go-cleanarch-starter/internal/infra/config"
)

func main() {
	// Load environment variables
	configLoad()

	// Connect to the database and close it when at its end
	db := databaseConn()
	defer db.Close()
}

func configLoad() {
	if err := config.LoadEnv(); err != nil {
		panic(err)
	}
}

// Open a new database connection
func databaseConn() *sql.DB {
	db, err := database.PostgreSQLConn()
	if err != nil {
		panic(err)
	}
	return db
}
