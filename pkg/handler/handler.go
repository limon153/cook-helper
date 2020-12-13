package handler

import (
	"cook-helper/pkg/service"
	"net/http"

	"github.com/go-chi/chi"
	"gopkg.in/go-playground/validator.v9"
)

// Handler container handlers for rest api
type Handler struct {
	validate *validator.Validate
	services *service.Service
}

// New create new Handler
func New(services *service.Service, validate *validator.Validate) *Handler {
	return &Handler{validate, services}
}

// InitRoutes initializes chi router and assign handlers to routes
func (h *Handler) InitRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello"))
	})

	router.Post("/sign-up", h.signUp)

	return router
}
