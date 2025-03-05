package usecase

import (
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/dto"
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/repository"
)

type FindSongsByArtistInternalUseCase struct {
	songRepository repository.SongRepository
}

func NewFindSongsByArtistUseCase(songRepository repository.SongRepository) *FindSongsByArtistInternalUseCase {
	return &FindSongsByArtistInternalUseCase{
		songRepository: songRepository,
	}
}

func (u *FindSongsByArtistInternalUseCase) Execute(artistID int) ([]*dto.SongDTO, error) {
	songs, err := u.songRepository.FindByArtist(artistID)
	if err != nil {
		return nil, err
	}

	var songsDTO []*dto.SongDTO
	for _, song := range songs {
		songsDTO = append(songsDTO, &dto.SongDTO{
			Title:       song.Title,
			Duration:    song.Duration,
			TrackNumber: song.TrackNumber,
		})
	}

	return songsDTO, nil
}
