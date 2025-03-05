package entity

// Song represents a music song with an ID, title, duration, track number.
// It depends on an Album.
type Song struct {
	ID          int
	Title       string
	Duration    int
	TrackNumber int
	Album       *Album
}

// NewSong creates a Song instance with the given ID, title, duration, and track number.
func NewSong(id int, title string, duration, trackNumber int) *Song {
	return &Song{
		ID:          id,
		Title:       title,
		Duration:    duration,
		TrackNumber: trackNumber,
	}
}
