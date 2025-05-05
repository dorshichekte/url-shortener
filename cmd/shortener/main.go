package main

import (
	"log"

	"url-shortener/internal/app/common"
	"url-shortener/internal/app/config"
	"url-shortener/internal/app/handlers"
	"url-shortener/internal/app/logger"
	"url-shortener/internal/app/server"
	"url-shortener/internal/app/services"
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
	store := storage.NewStorage(cfg, l)
	dependency := common.BaseDependency{
		Cfg:    *cfg,
		Logger: l,
	}
	service := services.NewServices(store, dependency)
	handler := handlers.NewHandlers(service, dependency)
	server.Start(cfg, handler, l)
}
