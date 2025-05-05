package handler

import (
	"url-shortener/internal/app/common"
	"url-shortener/internal/app/handler/db"
	"url-shortener/internal/app/handler/url"
	"url-shortener/internal/app/service"
)

type Handlers struct {
	URL      url.Method
	Database db.Method
}

type Handler struct {
	Handlers *Handlers
}

func NewHandlers(services service.Services, dependency common.BaseDependency) *Handler {
	handlers := &Handlers{
		URL:      url.NewURL(services, dependency),
		Database: db.NewDB(services, dependency),
	}

	return &Handler{
		Handlers: handlers,
	}
}
