package server

import (
	"net/http"

	"url-shortener/internal/app/handlers"
)

func Start() {
	mux := handlers.Register()

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		panic(err)
	}
}
