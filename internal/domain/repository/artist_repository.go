package repository

import "github.com/guilherme-or/go-cleanarch-starter/internal/domain/entity"

// ArtistRepository defines the interface for artist repository operations.
// It provides methods to find artists by various criteria.
type ArtistRepository interface {
	// FindAll retrieves all artists.
	// Returns a slice of pointers to Artist entities and an error if any.
	FindAll() ([]*entity.Artist, error)

	// FindByName retrieves an artist by their name.
	// Takes the artist's name as a parameter.
	// Returns a pointer to the Artist entity and an error if any.
	FindByName(name string) (*entity.Artist, error)

	// FindByNationality retrieves artists by their nationality.
	// Takes the artist's nationality as a parameter.
	// Returns a slice of pointers to Artist entities and an error if any.
	FindByNationality(nationality string) ([]*entity.Artist, error)
}
