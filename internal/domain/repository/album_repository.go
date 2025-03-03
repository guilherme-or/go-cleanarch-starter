package repository

import "github.com/guilherme-or/go-cleanarch-starter/internal/domain/entity"

type AlbumRepository interface {
	FindByYear(year int) ([]*entity.Album, error)
	FindByArtist(artistID int) ([]*entity.Album, error)
}
