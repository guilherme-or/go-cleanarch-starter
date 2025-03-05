package dto

import "time"

type AlbumDTO struct {
	Title       string     `json:"title"`
	ReleaseDate time.Time  `json:"release_date"`
	Artist      *ArtistDTO `json:"artist,omitempty"`
	Songs       []*SongDTO `json:"songs,omitempty"`
}
