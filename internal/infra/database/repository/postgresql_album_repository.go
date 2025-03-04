package database_repository

import (
	"database/sql"

	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/entity"
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/repository"
)

type PostgreSQLAlbumRepository struct {
	db *sql.DB
}

func NewPostgreSQLAlbumRepository(db *sql.DB) repository.AlbumRepository {
	return &PostgreSQLAlbumRepository{db: db}
}

// FindByArtist implements repository.AlbumRepository.
func (p *PostgreSQLAlbumRepository) FindByArtist(artistID int) ([]*entity.Album, error) {
	panic("unimplemented")
}

// FindByYear implements repository.AlbumRepository.
func (p *PostgreSQLAlbumRepository) FindByYear(year int) ([]*entity.Album, error) {
	panic("unimplemented")
}
