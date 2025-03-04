// Makefile run command entry point
// This file is used to build and run the entire application
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/guilherme-or/go-cleanarch-starter/internal/infra/database"
	database_repository "github.com/guilherme-or/go-cleanarch-starter/internal/infra/database/repository"
	artist_usecase "github.com/guilherme-or/go-cleanarch-starter/internal/usecase/artist"

	"github.com/guilherme-or/go-cleanarch-starter/internal/infra/config"
)

func main() {
	// Load environment
	cfg()

	// Establish database connection
	db := db()
	defer db.Close()

	// Instantiate dependencies
	deps(db)
}

// Loads environment configuration variables
func cfg() {
	if err := config.LoadEnv(); err != nil {
		panic(err)
	}
}

// Open a new database connection
func db() *sql.DB {
	db, err := database.PostgreSQLConn()
	if err != nil {
		panic(err)
	}
	return db
}

func deps(db *sql.DB) {
	// Repositories
	artistRepo := database_repository.NewPostgreSQLArtistRepository(db)
	// albumRepo := database_repository.NewPostgreSQLAlbumRepository(db)
	// songRepo := database_repository.NewPostgreSQLSongRepository(db)

	// Use Cases
	// Artist
	findAllArtistsUC := artist_usecase.NewFindAllArtistsUsecase(artistRepo)
	findArtistByNameUC := artist_usecase.NewFindArtistByNameUsecase(artistRepo)
	findArtistsByNationalityUC := artist_usecase.NewFindArtistsByNationalityUsecase(artistRepo)

	as, _ := findAllArtistsUC.Execute()
	j, _ := json.MarshalIndent(as, "", "  ")
	fmt.Println(string(j))

	a, _ := findArtistByNameUC.Execute("Daft Punk")
	j, _ = json.MarshalIndent(a, "", "  ")
	fmt.Println(string(j))

	as, _ = findArtistsByNationalityUC.Execute("American")
	j, _ = json.MarshalIndent(as, "", "  ")
	fmt.Println(string(j))
}
