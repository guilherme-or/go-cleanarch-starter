package entity

import "time"

// Artist represents a musical artist with an ID, name, birth date, and nationality.
// It can have multiple Albums.
type Artist struct {
	ID          int
	Name        string
	BirthDate   time.Time
	Nationality string
	Albums      []*Album
}

// NewArtist creates a Artist instance with the given ID, name, birth date, and nationality.
func NewArtist(id int, name string, birthDate time.Time, nationality string) *Artist {
	return &Artist{
		ID:          id,
		Name:        name,
		BirthDate:   birthDate,
		Nationality: nationality,
	}
}
