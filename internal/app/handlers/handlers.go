package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Register() http.Handler {
	r := chi.NewRouter()

	r.Post("/", Add)
	r.Get("/{id}", Get)

	return r
}
