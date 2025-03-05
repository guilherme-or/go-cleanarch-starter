// Makefile run command entry point
// This file is used to build and run the entire application
package main

import (
	"database/sql"

	"github.com/guilherme-or/go-cleanarch-starter/internal/infra/database"
	database_repository "github.com/guilherme-or/go-cleanarch-starter/internal/infra/database/repository"
	album_usecase "github.com/guilherme-or/go-cleanarch-starter/internal/usecase/album"
	artist_usecase "github.com/guilherme-or/go-cleanarch-starter/internal/usecase/artist"
	song_usecase "github.com/guilherme-or/go-cleanarch-starter/internal/usecase/song"

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
	albumRepo := database_repository.NewPostgreSQLAlbumRepository(db)
	songRepo := database_repository.NewPostgreSQLSongRepository(db)

	// Use Cases
	// Artist
	findAllArtistsUC := artist_usecase.NewFindAllArtistsUsecase(artistRepo)
	findArtistByNameUC := artist_usecase.NewFindArtistByNameUsecase(artistRepo)
	findArtistsByNationalityUC := artist_usecase.NewFindArtistsByNationalityUsecase(artistRepo)

	// Album
	findAlbumByTitleUC := album_usecase.NewFindAlbumByTitleUseCase(albumRepo)
	findAlbumsByArtistUC := album_usecase.NewFindAlbumsByArtistUseCase(albumRepo)
	findAlbumsByYearUC := album_usecase.NewFindAlbumsByYearUseCase(albumRepo)

	// Song
	findSongsByAlbumUC := song_usecase.NewFindSongsByAlbumUseCase(songRepo)
	findSongsByDurationIntervalUC := song_usecase.NewFindSongsByDurationInternalUseCase(songRepo)
	findSongsByArtistUC := song_usecase.NewFindSongsByArtistUseCase(songRepo)
}
