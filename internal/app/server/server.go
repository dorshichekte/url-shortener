package server

import (
	"log"
	"net/http"

	"url-shortener/internal/app/config"
	"url-shortener/internal/app/handlers"
)

func Start(cfg *config.Config, handler *handlers.Handler) {
	mux := handler.Register()
	log.Fatal(http.ListenAndServe(cfg.ServerAddress, mux))
}
