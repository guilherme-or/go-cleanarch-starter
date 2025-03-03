package entity

import "time"

// Album represents a music album with an ID, title, release date, artist.
// It depends on an Artist and can have multiple Songs.
type Album struct {
	ID          int
	Title       string
	ReleaseDate time.Time
	Artist      *Artist
	Songs       []*Song
}

// NewAlbum creates an Album instance with the given ID, title, and release date.
func NewAlbum(id int, title string, releaseDate time.Time) *Album {
	return &Album{
		ID:          id,
		Title:       title,
		ReleaseDate: releaseDate,
	}
}
