package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	adapter "url-shortener/internal/app/config/adapter"
)

type Router struct {
	router *chi.Mux
	logger *zap.Logger
	config *adapter.Router
}

type Route struct {
	Method  string
	Path    string
	Handler http.Handler
}
