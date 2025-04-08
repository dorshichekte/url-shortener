package handlers

import (
	"go.uber.org/zap"
	"net/http"
	"url-shortener/internal/app/middleware"

	"github.com/go-chi/chi/v5"
	"url-shortener/internal/app/config"
	"url-shortener/internal/app/handlers/url"
	u "url-shortener/internal/app/services/url"
)

func NewHandler(urlService *u.Service, cfg *config.Config, logger *zap.Logger) *Handler {
	return &Handler{
		urlHandler: url.NewHandler(urlService, cfg, logger),
	}
}

func (h *Handler) Register(logger *zap.Logger) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Log(logger))
	r.Use(middleware.Gzip)
	r.Use(middleware.Decompress)

	r.Post("/", h.urlHandler.Add)
	r.Get("/{id}", h.urlHandler.Get)
	r.Post("/api/shorten", h.urlHandler.Shorten)

	return r
}
