package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/dto"
	usecase "github.com/guilherme-or/go-cleanarch-starter/internal/usecase/album"
)

type AlbumController struct {
	findAlbumByTitleUC   *usecase.FindAlbumByTitleUseCase
	findAlbumsByArtistUC *usecase.FindAlbumsByArtistUseCase
	findAlbumsByYearUC   *usecase.FindAlbumsByYearUseCase
}

func NewAlbumController(
	findAlbumByTitleUC *usecase.FindAlbumByTitleUseCase,
	findAlbumsByArtistUC *usecase.FindAlbumsByArtistUseCase,
	findAlbumsByYearUC *usecase.FindAlbumsByYearUseCase,
) *AlbumController {
	return &AlbumController{
		findAlbumByTitleUC:   findAlbumByTitleUC,
		findAlbumsByArtistUC: findAlbumsByArtistUC,
		findAlbumsByYearUC:   findAlbumsByYearUC,
	}
}

func (c *AlbumController) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/albums", c.FindAlbums)
	router.HandleFunc("/albums/{title}", c.FindAlbumByTitle)
}

func (c *AlbumController) FindAlbums(w http.ResponseWriter, r *http.Request) {
	var albums []*dto.AlbumDTO
	var err error
	year := r.URL.Query().Get("year")
	artistID := r.URL.Query().Get("artist_id")

	if year == "" && artistID == "" {
		http.Error(
			w,
			"Year or Artist ID required on query parameter (?year= or ?artist_id=)",
			http.StatusBadRequest,
		)
		return
	} else if year != "" && artistID == "" {
		year, convErr := strconv.Atoi(year)
		if convErr != nil {
			http.Error(w, convErr.Error(), http.StatusBadRequest)
			return
		}
		albums, err = c.findAlbumsByYearUC.Execute(year)
	} else if (year == "" && artistID != "") || (year != "" && artistID != "") {
		// Even if both are present, prioritize artist ID
		id, convErr := strconv.Atoi(artistID)
		if convErr != nil {
			http.Error(w, convErr.Error(), http.StatusBadRequest)
			return
		}
		albums, err = c.findAlbumsByArtistUC.Execute(id)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	m, err := json.Marshal(albums)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(m)
}

func (c *AlbumController) FindAlbumByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.PathValue("title")
	album, err := c.findAlbumByTitleUC.Execute(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	m, err := json.Marshal(album)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(m)
}
