// Makefile run command entry point
// This file is used to build and run the entire application
package main

import (
	"database/sql"
	"github.com/guilherme-or/go-cleanarch-starter/internal/infra/database"

	"github.com/guilherme-or/go-cleanarch-starter/internal/config"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to the database and close it when at its end
	db := databaseConn()
	defer db.Close()
}

// Open a new database connection
func databaseConn() *sql.DB {
	db, err := database.PostgreSQLConn()
	if err != nil {
		panic(err)
	}
	return db
}
