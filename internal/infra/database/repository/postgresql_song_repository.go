package database_repository

import (
	"database/sql"

	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/entity"
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/repository"
)

type PostgreSQLSongRepository struct {
	db *sql.DB
}

func NewPostgreSQLSongRepository(db *sql.DB) repository.SongRepository {
	return &PostgreSQLSongRepository{db: db}
}

// FindByAlbum implements repository.SongRepository.
func (p *PostgreSQLSongRepository) FindByAlbum(albumID int) ([]*entity.Song, error) {
	panic("unimplemented")
}

// FindByArtist implements repository.SongRepository.
func (p *PostgreSQLSongRepository) FindByArtist(artistID int) ([]*entity.Song, error) {
	panic("unimplemented")
}

// FindByDuration implements repository.SongRepository.
func (p *PostgreSQLSongRepository) FindByDuration(durationInS int) ([]*entity.Song, error) {
	panic("unimplemented")
}
