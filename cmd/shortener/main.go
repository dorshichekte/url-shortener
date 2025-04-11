package main

import (
	"log"

	"url-shortener/internal/app/config"
	"url-shortener/internal/app/handlers"
	"url-shortener/internal/app/logger"
	"url-shortener/internal/app/server"
	"url-shortener/internal/app/services/url"
	"url-shortener/internal/app/storage"
)

func main() {
	l, err := logger.New()
	if err != nil {
		log.Fatalf("Failed initialization logger: %v", err)
	}
	defer func() {
		_ = l.Sync()
	}()

	cfg := config.NewConfig()

	urlStorage := storage.Create(cfg, l)

	urlService := url.NewURLService(urlStorage, cfg)

	handler := handlers.NewHandler(urlService, cfg, l)

	server.Start(cfg, handler, l)
}
