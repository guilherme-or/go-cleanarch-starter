package usecase

import (
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/dto"
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/repository"
)

type FindAlbumsByYearUseCase struct {
	albumRepository repository.AlbumRepository
}

func NewFindAlbumByYearUseCase(albumRepository repository.AlbumRepository) *FindAlbumsByYearUseCase {
	return &FindAlbumsByYearUseCase{
		albumRepository: albumRepository,
	}
}

func (u *FindAlbumsByYearUseCase) Execute(year int) ([]*dto.AlbumDTO, error) {
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
