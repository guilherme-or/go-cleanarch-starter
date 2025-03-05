package entity_test

import (
	"testing"

	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/entity"
)

var (
	songId = 1
	songTitle = "Test Song"
	songDuration = 300
	songTrackNumber = 1
)

func TestNewSong(t *testing.T) {
	song := entity.NewSong(songId, songTitle, songDuration, songTrackNumber)

	if song.ID != songId {
		t.Errorf("expected ID %d, got %d", songId, song.ID)
	}
	if song.Title != songTitle {
		t.Errorf("expected Title %s, got %s", songTitle, song.Title)
	}
	if song.Duration != songDuration {
		t.Errorf("expected Duration %d, got %d", songDuration, song.Duration)
	}
	if song.TrackNumber != songTrackNumber {
		t.Errorf("expected TrackNumber %d, got %d", songTrackNumber, song.TrackNumber)
	}
	if song.Album != nil {
		t.Errorf("expected Album to be nil, got %v", song.Album)
	}
}
