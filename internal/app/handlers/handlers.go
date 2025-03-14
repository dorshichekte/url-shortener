package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"url-shortener/internal/app/config"
	"url-shortener/internal/app/handlers/url"
	u "url-shortener/internal/app/services/url"
)

func NewHandler(urlService *u.Service, cfg *config.Config) *Handler {
	return &Handler{
		urlHandler: url.NewHandler(urlService, cfg),
	}
}

func (h *Handler) Register() http.Handler {
	r := chi.NewRouter()

	r.Post("/", h.urlHandler.Add)
	r.Get("/{id}", h.urlHandler.Get)

	return r
}
