package repository

import "github.com/guilherme-or/go-cleanarch-starter/internal/domain/entity"

// SongRepository defines methods to find songs based on various criteria.
type SongRepository interface {
	FindByDuration(startDuration, endDuration int) ([]*entity.Song, error) // Retrieves a list of songs that match the specified duration range in seconds.
	FindByAlbum(albumID int) ([]*entity.Song, error)                       // Retrieves a list of songs associated with the specified album ID.
	FindByArtist(artistID int) ([]*entity.Song, error)                     // Retrieves a list of songs associated with the specified artist ID.
}
