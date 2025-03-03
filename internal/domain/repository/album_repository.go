package repository

import "github.com/guilherme-or/go-cleanarch-starter/internal/domain/entity"

// AlbumRepository defines the methods that any implementation of an album repository should have.
// It provides methods to retrieve albums based on specific criteria such as release year and artist ID.
type AlbumRepository interface {
	FindByYear(year int) ([]*entity.Album, error)       // Retrieves a list of albums released in the specified year.
	FindByArtist(artistID int) ([]*entity.Album, error) // Retrieves a list of albums by the specified artist.
}
