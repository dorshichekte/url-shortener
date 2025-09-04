package httpserver

import (
	"net/http"
	"url-shortener/internal/app/config"

	"go.uber.org/zap"
)

// Http инкапсулирует HTTP-сервер и его зависимости.
type Http struct {
	logger *zap.Logger
	server *http.Server
	config *config.Config
}
