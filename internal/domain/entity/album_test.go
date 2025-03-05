package entity_test

import (
	"testing"
	"time"

	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/entity"
)

var (
	albumId          = 1
	albumTitle       = "Test Album"
	albumReleaseDate = time.Now()
)

func TestNewAlbum(t *testing.T) {
	album := entity.NewAlbum(albumId, albumTitle, albumReleaseDate)

	if album == nil {
		t.Fatalf("Expected album to be non-nil")
	}
	if album.ID != albumId {
		t.Errorf("Expected album ID to be %d, got %d", albumId, album.ID)
	}
	if album.Title != albumTitle {
		t.Errorf("Expected album title to be %s, got %s", albumTitle, album.Title)
	}
	if !album.ReleaseDate.Equal(albumReleaseDate) {
		t.Errorf("Expected album release date to be %v, got %v", albumReleaseDate, album.ReleaseDate)
	}
	if album.Artist != nil {
		t.Errorf("Expected album artist to be nil, got %v", album.Artist)
	}
	if album.Songs != nil {
		t.Errorf("Expected album tracks to be nil, got %v", album.Songs)
	}
}
