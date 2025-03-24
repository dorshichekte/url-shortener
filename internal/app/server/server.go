package server

import (
	"go.uber.org/zap"
	"log"
	"net/http"

	"url-shortener/internal/app/config"
	"url-shortener/internal/app/handlers"
)

func Start(cfg *config.Config, handler *handlers.Handler, logger *zap.Logger) {
	mux := handler.Register(logger)
	log.Fatal(http.ListenAndServe(cfg.ServerAddress, mux))
}
