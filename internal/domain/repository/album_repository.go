package repository

import "github.com/guilherme-or/go-cleanarch-starter/internal/domain/entity"

// AlbumRepository defines methods to retrieve albums based on specific criteria.
type AlbumRepository interface {
	FindByYear(year int) ([]*entity.Album, error)       // Retrieves a list of albums released in the specified year.
	FindByArtist(artistID int) ([]*entity.Album, error) // Retrieves a list of albums by the specified artist.
}
