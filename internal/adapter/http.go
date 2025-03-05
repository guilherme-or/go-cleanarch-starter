package adapter

import "net/http"

type Server interface {
	Start()
	GetRouter() *http.ServeMux
}

type Controller interface {
	RegisterRoutes(router *http.ServeMux)
}
