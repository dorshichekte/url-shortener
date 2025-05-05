package server

import (
	"go.uber.org/zap"
	"net/http"
	"url-shortener/internal/app/constants"

	"url-shortener/internal/app/config"
	"url-shortener/internal/app/handlers"
)

func Start(cfg *config.AppConfig, handler *handlers.Handler, logger *zap.Logger) {
	mux := handler.Register(logger)
	
	err := http.ListenAndServe(cfg.ServerAddress, mux)
	if err != nil {
		logger.Fatal(constants.ErrServerDown.Error(), zap.Error(err))
	}
}
