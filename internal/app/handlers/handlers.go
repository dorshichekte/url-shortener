package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Register() http.Handler {
	r := chi.NewRouter()

	r.Post("/", AddURL)
	r.Get("/{id}", GetURL)

	return r
}
