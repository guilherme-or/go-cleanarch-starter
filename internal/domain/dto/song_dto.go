package dto

type SongDTO struct {
	Title       string    `json:"title"`
	Duration    int       `json:"duration"`
	TrackNumber int       `json:"track_number"`
	Album       *AlbumDTO `json:"album,omitempty"`
}
