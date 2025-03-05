package usecase

import (
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/dto"
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/repository"
)

type FindAlbumsByArtistUseCase struct {
	albumRepository repository.AlbumRepository
}

func NewFindAlbumByArtistUseCase(albumRepository repository.AlbumRepository) *FindAlbumsByArtistUseCase {
	return &FindAlbumsByArtistUseCase{
		albumRepository: albumRepository,
	}
}

func (u *FindAlbumsByArtistUseCase) Execute(artistID int) ([]*dto.AlbumDTO, error) {
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
