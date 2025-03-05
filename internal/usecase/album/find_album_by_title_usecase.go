package usecase

import (
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/dto"
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/repository"
)

type FindAlbumByTitleUseCase struct {
	albumRepository repository.AlbumRepository
}

func NewFindAlbumByTitleUseCase(albumRepository repository.AlbumRepository) *FindAlbumByTitleUseCase {
	return &FindAlbumByTitleUseCase{
		albumRepository: albumRepository,
	}
}

func (u *FindAlbumByTitleUseCase) Execute(title string) (*dto.AlbumDTO, error) {
	album, err := u.albumRepository.FindByTitle(title)
	if err != nil {
		return nil, err
	}

	albumDTO := &dto.AlbumDTO{
		Title:       album.Title,
		ReleaseDate: album.ReleaseDate,
		Artist: &dto.ArtistDTO{
			Name:        album.Artist.Name,
			BirthDate:   album.Artist.BirthDate,
			Nationality: album.Artist.Nationality,
		},
	}

	var songsDTO []*dto.SongDTO
	for _, song := range album.Songs {
		songsDTO = append(songsDTO, &dto.SongDTO{
			Title:       song.Title,
			Duration:    song.Duration,
			TrackNumber: song.TrackNumber,
		})
	}

	albumDTO.Songs = songsDTO

	return albumDTO, nil
}
