package server

import (
	"net/http"

	"url-shortener/internal/app/config"
	"url-shortener/internal/app/handlers"
)

func Start() {
	mux := handlers.Register()
	conf := config.GetConfig()
	as := conf.ServerAddress.String()

	err := http.ListenAndServe(as, mux)

	if err != nil {
		panic(err)
	}
}
