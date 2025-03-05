package entity_test

import (
	"testing"
	"time"

	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/entity"
)

var (
	artistId          = 1
	artistName        = "John Doe"
	artistBirthDate   = time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)
	artistNationality = "American"
)

// NewArtist function should return an pointer to a Artist struct
// with the given ID, name, birthDate and nationality.
func TestNewArtist(t *testing.T) {
	artist := entity.NewArtist(artistId, artistName, artistBirthDate, artistNationality)

	if artist == nil {
		t.Fatal("expected Artist instance, got nil")
	}
	if artist.ID != artistId {
		t.Errorf("expected ID %d, got %d", artistId, artist.ID)
	}
	if artist.Name != artistName {
		t.Errorf("expected Name %s, got %s", artistName, artist.Name)
	}
	if artist.BirthDate != artistBirthDate {
		t.Errorf("expected BirthDate %v, got %v", artistBirthDate, artist.BirthDate)
	}
	if artist.Nationality != artistNationality {
		t.Errorf("expected Nationality %s, got %s", artistNationality, artist.Nationality)
	}
	if artist.Albums != nil {
		t.Errorf("expected Albums to be nil, got %v", artist.Albums)
	}
}
