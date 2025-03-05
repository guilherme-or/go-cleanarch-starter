package usecase

import (
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/dto"
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/repository"
)

type FindAllArtistsUseCase struct {
	artistRepository repository.ArtistRepository
}

func NewFindAllArtistsUsecase(artistRepository repository.ArtistRepository) *FindAllArtistsUseCase {
	return &FindAllArtistsUseCase{
		artistRepository: artistRepository,
	}
}

func (u *FindAllArtistsUseCase) Execute() ([]*dto.ArtistDTO, error) {
	artists, err := u.artistRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var artistDTOs []*dto.ArtistDTO
	for _, artist := range artists {
		artistDTOs = append(artistDTOs, &dto.ArtistDTO{
			Name:        artist.Name,
			BirthDate:   artist.BirthDate,
			Nationality: artist.Nationality,
		})
	}

	return artistDTOs, nil
}
