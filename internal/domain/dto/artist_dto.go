package dto

import "time"

type ArtistDTO struct {
	Name        string      `json:"name"`
	BirthDate   time.Time   `json:"birth_date"`
	Nationality string      `json:"nationality"`
	Albums      []*AlbumDTO `json:"albums,omitempty"`
}
