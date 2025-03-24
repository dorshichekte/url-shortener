package main

import (
	"go.uber.org/zap"
	"log"

	"url-shortener/internal/app/config"
	"url-shortener/internal/app/handlers"
	"url-shortener/internal/app/logger"
	"url-shortener/internal/app/server"
	"url-shortener/internal/app/services/url"
	"url-shortener/internal/app/storage"
)

func main() {
	logger, err := logger.New()
	if err != nil {
		log.Fatalf("Failed initialization logger: %v", err)
	}
	defer logger.Sync()

	cfg := config.NewConfig()

	urlStorage := storage.NewURLStorage()

	urlService := url.NewURLService(urlStorage)

	handler := handlers.NewHandler(urlService, cfg)

	server.Start(cfg, handler, logger)
}
