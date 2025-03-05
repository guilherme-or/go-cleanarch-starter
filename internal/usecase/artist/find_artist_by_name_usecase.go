package usecase

import (
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/dto"
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/repository"
)

type FindArtistByNameUseCase struct {
	artistRepository repository.ArtistRepository
}

func NewFindArtistByNameUsecase(artistRepository repository.ArtistRepository) *FindArtistByNameUseCase {
	return &FindArtistByNameUseCase{
		artistRepository: artistRepository,
	}
}

func (u *FindArtistByNameUseCase) Execute(name string) (*dto.ArtistDTO, error) {
	artist, err := u.artistRepository.FindByName(name)
	if err != nil {
		return nil, err
	}

	artistDTO := &dto.ArtistDTO{
		Name:        artist.Name,
		BirthDate:   artist.BirthDate,
		Nationality: artist.Nationality,
	}

	artistDTO.Albums = make([]*dto.AlbumDTO, 0)
	if artist.Albums != nil && len(artist.Albums) > 0 {
		for _, album := range artist.Albums {
			artistDTO.Albums = append(artistDTO.Albums, &dto.AlbumDTO{
				Title: album.Title,
				ReleaseDate: album.ReleaseDate,
			})
		}
	}

	
	return artistDTO, nil
}
