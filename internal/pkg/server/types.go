package server

import (
	"net/http"

	"go.uber.org/zap"

	config "url-shortener/internal/app/config/adapter"
)

type Server struct {
	logger *zap.Logger
	server *http.Server
	config *config.HTTPServer
}
