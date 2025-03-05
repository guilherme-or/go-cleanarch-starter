package usecase

import (
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/dto"
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/repository"
)

type FindAlbumsByYearUsecase struct {
	albumRepository repository.AlbumRepository
}

func NewFindAlbumByYearUsecase(albumRepository repository.AlbumRepository) *FindAlbumsByYearUsecase {
	return &FindAlbumsByYearUsecase{
		albumRepository: albumRepository,
	}
}

func (u *FindAlbumsByYearUsecase) Execute(year int) ([]*dto.AlbumDTO, error) {
	albums, err := u.albumRepository.FindByYear(year)
	if err != nil {
		return nil, err
	}

	var albumsDTO []*dto.AlbumDTO
	for _, album := range albums {
		albumsDTO = append(albumsDTO, &dto.AlbumDTO{
			Title:       album.Title,
			ReleaseDate: album.ReleaseDate,
		})
	}

	return albumsDTO, nil
}
