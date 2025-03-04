package repository

import "github.com/guilherme-or/go-cleanarch-starter/internal/domain/entity"

// ArtistRepository defines methods to find artists by various criteria.
type ArtistRepository interface {
	FindAll() ([]*entity.Artist, error)                             // Retrieves all artists. Fails if none are found.
	FindByName(name string) (*entity.Artist, error)                 // Retrieves an artist by their name. Fails if none are found.
	FindByNationality(nationality string) ([]*entity.Artist, error) // Retrieves artists by their nationality. Fails if none are found.
}
