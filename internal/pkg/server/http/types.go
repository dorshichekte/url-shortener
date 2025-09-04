package httpserver

import (
	"net/http"
	"url-shortener/internal/app/config"

	"go.uber.org/zap"
)

// HTTP инкапсулирует HTTP-сервер и его зависимости.
type HTTP struct {
	logger *zap.Logger
	server *http.Server
	config *config.Config
}
