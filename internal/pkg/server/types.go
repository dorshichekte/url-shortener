package server

import (
	"net/http"

	"go.uber.org/zap"

	config "url-shortener/internal/app/config/adapter"
)

// Server инкапсулирует HTTP-сервер и его зависимости.
type Server struct {
	logger *zap.Logger
	server *http.Server
	config *config.HTTPServer
}
