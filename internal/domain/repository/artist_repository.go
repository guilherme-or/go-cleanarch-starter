package repository

import "github.com/guilherme-or/go-cleanarch-starter/internal/domain/entity"

type ArtistRepository interface {
	FindAll() ([]*entity.Artist, error)
	FindByName(name string) (*entity.Artist, error)
	FindByNationality(nationality string) ([]*entity.Artist, error)
}
