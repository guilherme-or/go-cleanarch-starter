package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/dto"
	usecase "github.com/guilherme-or/go-cleanarch-starter/internal/usecase/song"
)

type SongController struct {
	findSongsByDurationIntervalUC *usecase.FindSongsByDurationIntervalUseCase
	findSongsByAlbumUC            *usecase.FindSongsByAlbumUseCase
	findSongsByArtistUC           *usecase.FindSongsByArtistUseCase
}

func NewSongController(
	findSongsByDurationIntervalUC *usecase.FindSongsByDurationIntervalUseCase,
	findSongsByAlbumUC *usecase.FindSongsByAlbumUseCase,
	findSongsByArtistUC *usecase.FindSongsByArtistUseCase,
) *SongController {
	return &SongController{
		findSongsByDurationIntervalUC: findSongsByDurationIntervalUC,
		findSongsByAlbumUC:            findSongsByAlbumUC,
		findSongsByArtistUC:           findSongsByArtistUC,
	}
}

func (c *SongController) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/songs", c.FindSongs)
	router.HandleFunc("/songs/duration", c.FindSongsByDurationInterval)
}

func (c *SongController) FindSongs(w http.ResponseWriter, r *http.Request) {
	var songs []*dto.SongDTO
	var err error
	albumID := r.URL.Query().Get("album_id")
	artistID := r.URL.Query().Get("artist_id")

	if albumID == "" && artistID == "" {
		http.Error(
			w,
			"Album ID or Artist ID required on query parameter (?album_id= or ?artist_id=)",
			http.StatusBadRequest,
		)
		return
	} else if artistID != "" && albumID == "" {
		id, convErr := strconv.Atoi(artistID)
		if convErr != nil {
			http.Error(w, convErr.Error(), http.StatusBadRequest)
			return
		}
		songs, err = c.findSongsByArtistUC.Execute(id)
	} else if (artistID == "" && albumID != "") || (artistID != "" && albumID != "") {
		// Even if both are present, prioritize album ID
		id, convErr := strconv.Atoi(albumID)
		if convErr != nil {
			http.Error(w, convErr.Error(), http.StatusBadRequest)
			return
		}
		songs, err = c.findSongsByAlbumUC.Execute(id)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	m, err := json.Marshal(songs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(m)
}

func (c *SongController) FindSongsByDurationInterval(w http.ResponseWriter, r *http.Request) {
	var songs []*dto.SongDTO
	var err error
	startDuration := r.URL.Query().Get("start")
	endDuration := r.URL.Query().Get("end")

	if startDuration == "" || endDuration == "" {
		http.Error(
			w,
			"Song start and end duration required on query parameter (?start= and ?end=)",
			http.StatusBadRequest,
		)
		return
	}

	start, convErr := strconv.Atoi(startDuration)
	if convErr != nil {
		http.Error(w, convErr.Error(), http.StatusBadRequest)
		return
	}
	end, convErr := strconv.Atoi(endDuration)
	if convErr != nil {
		http.Error(w, convErr.Error(), http.StatusBadRequest)
		return
	}

	songs, err = c.findSongsByDurationIntervalUC.Execute(start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	m, err := json.Marshal(songs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(m)
}
