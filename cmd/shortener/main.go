package main

import (
	"log"

	"url-shortener/internal/app/common"
	"url-shortener/internal/app/config"
	"url-shortener/internal/app/handler"
	"url-shortener/internal/app/logger"
	"url-shortener/internal/app/server"
	"url-shortener/internal/app/service"
	"url-shortener/internal/app/storage"
)

func main() {
	l, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("Failed initialization logger: %v", err)
	}
	defer func() {
		_ = l.Sync()
	}()

	cfg := config.NewConfig(l)
	store := storage.NewStorage(cfg.App, l)
	dependency := common.BaseDependency{
		Cfg:    *cfg.App,
		Logger: l,
	}
	service := service.NewServices(store, dependency)
	defer service.Worker.Close()

	handler := handler.NewHandlers(service, dependency)
	server.Start(cfg.App, handler, l)
}
