package usecase

import (
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/dto"
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/repository"
)

type FindSongsByDurationInternalUseCase struct {
	songRepository repository.SongRepository
}

func NewFindSongsByDurationInternalUseCase(songRepository repository.SongRepository) *FindSongsByDurationInternalUseCase {
	return &FindSongsByDurationInternalUseCase{
		songRepository: songRepository,
	}
}

func (u *FindSongsByDurationInternalUseCase) Execute(startDuration, endDuration int) ([]*dto.SongDTO, error) {
	songs, err := u.songRepository.FindByDuration(startDuration, endDuration)
	if err != nil {
		return nil, err
	}

	var songsDTO []*dto.SongDTO
	for _, song := range songs {
		songsDTO = append(songsDTO, &dto.SongDTO{
			Title:       song.Title,
			Duration:    song.Duration,
			TrackNumber: song.TrackNumber,
			Album: &dto.AlbumDTO{
				Title:       song.Album.Title,
				ReleaseDate: song.Album.ReleaseDate,
			},
		})
	}

	return songsDTO, nil
}
