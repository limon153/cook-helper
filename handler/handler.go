package handler

import (
	"cook-helper/service"
	"net/http"

	"github.com/go-chi/chi"
)

// Handler container handlers for rest api
type Handler struct {
	services *service.Service
}

// New create new Handler
func New(services *service.Service) *Handler {
	return &Handler{services: services}
}

// InitRoutes initializes chi router and assign handlers to routes
func (h *Handler) InitRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello"))
	})

	return router
}
