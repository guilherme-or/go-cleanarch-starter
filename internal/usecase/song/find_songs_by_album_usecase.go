package usecase

import (
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/dto"
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/repository"
)

type FindSongsByAlbumUseCase struct {
	songRepository repository.SongRepository
}

func NewFindSongsByAlbumUseCase(songRepository repository.SongRepository) *FindSongsByAlbumUseCase {
	return &FindSongsByAlbumUseCase{
		songRepository: songRepository,
	}
}

func (u *FindSongsByAlbumUseCase) Execute(albumID int) ([]*dto.SongDTO, error) {
	songs, err := u.songRepository.FindByAlbum(albumID)
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
