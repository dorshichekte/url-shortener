package server

import (
	"fmt"
	"net/http"

	"url-shortener/internal/app/config"
	"url-shortener/internal/app/handlers"
)

func Start() {
	mux := handlers.Register()
	cfg := config.Get()
	fmt.Println(cfg)
	err := http.ListenAndServe(cfg.ServerAddress, mux)

	if err != nil {
		panic(err)
	}
}
