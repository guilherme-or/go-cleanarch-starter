package entity_test

import (
	"testing"
	"time"

	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/entity"
)

var (
	id          = 1
	name        = "John Doe"
	birthDate   = time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)
	nationality = "American"
)

// NewArtist function should return an pointer to a Artist struct
// with the given ID, name, birthDate and nationality.
func TestNewArtist(t *testing.T) {
	artist := entity.NewArtist(id, name, birthDate, nationality)

	if artist == nil {
		t.Fatal("expected Artist instance, got nil")
	}
	if artist.ID != id {
		t.Errorf("expected ID %d, got %d", id, artist.ID)
	}
	if artist.Name != name {
		t.Errorf("expected Name %s, got %s", name, artist.Name)
	}
	if artist.BirthDate != birthDate {
		t.Errorf("expected BirthDate %v, got %v", birthDate, artist.BirthDate)
	}
	if artist.Nationality != nationality {
		t.Errorf("expected Nationality %s, got %s", nationality, artist.Nationality)
	}
	if artist.Albums != nil {
		t.Errorf("expected Albums to be nil, got %v", artist.Albums)
	}
}
