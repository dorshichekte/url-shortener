package main

import (
	"url-shortener/internal/app/config"
	"url-shortener/internal/app/handlers"
	"url-shortener/internal/app/server"
	"url-shortener/internal/app/services/url"
	"url-shortener/internal/app/storage"
)

func main() {
	cfg := config.NewConfig()

	urlStorage := storage.NewURLStorage()

	urlService := url.NewURLService(urlStorage)

	handler := handlers.NewHandler(urlService, cfg)

	server.Start(cfg, handler)
}
