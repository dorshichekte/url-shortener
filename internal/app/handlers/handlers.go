package handlers

import (
	"net/http"
	"url-shortener/internal/app/common"
	"url-shortener/internal/app/services"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"url-shortener/internal/app/handlers/db"
	"url-shortener/internal/app/handlers/url"
	"url-shortener/internal/app/middleware"
)

func NewHandlers(services services.Services, dependency common.BaseDependency) *Handler {
	handlers := &Handlers{
		URL:      url.New(services, dependency),
		Database: db.New(services, dependency),
	}

	return &Handler{
		Handlers: handlers,
	}
}

func (h *Handler) Register(logger *zap.Logger) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Log(logger))
	r.Use(middleware.Gzip)
	r.Use(middleware.Decompress)

	r.Get("/ping", h.Handlers.Database.Ping)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth)
		r.Post("/", h.Handlers.URL.Create)
		r.Get("/{id}", h.Handlers.URL.GetByID)
		r.Get("/api/user/urls", h.Handlers.URL.GetAllByUser)
		r.Post("/api/shorten/batch", h.Handlers.URL.CreateBatch)
		r.Post("/api/shorten", h.Handlers.URL.CreateFromJSON)
		r.Delete("/api/users/urls", h.Handlers.URL.GetAllByUser)
	})

	return r
}
