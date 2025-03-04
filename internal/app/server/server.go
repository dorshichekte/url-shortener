package server

import (
	"net/http"

	"url-shortener/internal/app/config"
	"url-shortener/internal/app/handlers"
)

func Start() {
	mux := handlers.Register()
	cfg := config.Get()

	err := http.ListenAndServe(cfg.ServerAddress, mux)

	if err != nil {
		panic(err)
	}
}
