package usecase

import (
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/dto"
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/repository"
)

type FindArtistsByNationalityUseCase struct {
	artistRepository repository.ArtistRepository
}

func NewFindArtistsByNationalityUsecase(artistRepository repository.ArtistRepository) *FindArtistsByNationalityUseCase {
	return &FindArtistsByNationalityUseCase{
		artistRepository: artistRepository,
	}
}

func (u *FindArtistsByNationalityUseCase) Execute(nationality string) ([]*dto.ArtistDTO, error) {
	artists, err := u.artistRepository.FindByNationality(nationality)
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
