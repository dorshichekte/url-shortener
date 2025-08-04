package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	adapter "url-shortener/internal/app/config/adapter"
)

// Router содержит настройки маршрутизации, логгер и конфигурацию.
type Router struct {
	router *chi.Mux
	logger *zap.Logger
	config *adapter.Router
}

// Route описывает отдельный HTTP-маршрут.
type Route struct {
	Method  string
	Path    string
	Handler http.Handler
}
