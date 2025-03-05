package api

import (
	"log"
	"net/http"

	"github.com/guilherme-or/go-cleanarch-starter/internal/adapter"
	"github.com/guilherme-or/go-cleanarch-starter/internal/infra/config"
)

type APIServer struct {
	router *http.ServeMux
	logger *log.Logger
}

func NewAPIServer(controllers []adapter.Controller) adapter.Server {
	router := http.NewServeMux()
	logger := log.Default()

	for _, controller := range controllers {
		controller.RegisterRoutes(router)
	}

	return &APIServer{
		router: router,
		logger: logger,
	}
}

// GetRouter implements adapter.Server.
func (a *APIServer) GetRouter() *http.ServeMux {
	return a.router
}

// Start implements adapter.Server.
func (a *APIServer) Start() {
	a.logger.Println("Server is running on", config.E.Server.Addr)
	log.Fatal(http.ListenAndServe(config.E.Server.Addr, a.router))
}
