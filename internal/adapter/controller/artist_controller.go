package controller

import (
	"encoding/json"
	"net/http"

	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/dto"
	usecase "github.com/guilherme-or/go-cleanarch-starter/internal/usecase/artist"
)

type ArtistController struct {
	findAllArtistsUC           *usecase.FindAllArtistsUseCase
	findArtistByNameUC         *usecase.FindArtistByNameUseCase
	findArtistsByNationalityUC *usecase.FindArtistsByNationalityUseCase
}

func NewArtistController(
	findAllArtistsUC *usecase.FindAllArtistsUseCase,
	findArtistByNameUC *usecase.FindArtistByNameUseCase,
	findArtistsByNationalityUC *usecase.FindArtistsByNationalityUseCase,
) *ArtistController {
	return &ArtistController{
		findAllArtistsUC:           findAllArtistsUC,
		findArtistByNameUC:         findArtistByNameUC,
		findArtistsByNationalityUC: findArtistsByNationalityUC,
	}
}

func (c *ArtistController) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/artists", c.FindArtists)
	router.HandleFunc("/artists/{name}", c.FindArtistByName)
}

func (c *ArtistController) FindArtists(w http.ResponseWriter, r *http.Request) {
	var artists []*dto.ArtistDTO
	var err error

	nationality := r.URL.Query().Get("nationality")
	if nationality == "" {
		artists, err = c.findAllArtistsUC.Execute()
	} else {
		artists, err = c.findArtistsByNationalityUC.Execute(nationality)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	m, err := json.Marshal(artists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(m)
}

func (c *ArtistController) FindArtistByName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	artist, err := c.findArtistByNameUC.Execute(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	m, err := json.Marshal(artist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(m)
}
