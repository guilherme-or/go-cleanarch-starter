package usecase

import (
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/dto"
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/repository"
)

type FindAlbumsByArtistUsecase struct {
	albumRepository repository.AlbumRepository
}

func NewFindAlbumByArtistUsecase(albumRepository repository.AlbumRepository) *FindAlbumsByArtistUsecase {
	return &FindAlbumsByArtistUsecase{
		albumRepository: albumRepository,
	}
}

func (u *FindAlbumsByArtistUsecase) Execute(artistID int) ([]*dto.AlbumDTO, error) {
	albums, err := u.albumRepository.FindByArtist(artistID)
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
