package server

import (
	"net/http"

	"go.uber.org/zap"

	"url-shortener/internal/app/config"
)

// Server инкапсулирует HTTP-сервер и его зависимости.
type Server struct {
	logger *zap.Logger
	server *http.Server
	config *config.Config
}
