package handlers

import (
	"net/http"

	"url-shortener/internal/app/handlers/middlewares"
)

func Register() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", AddURL)
	mux.HandleFunc("/{id}", GetURL)

	wrappedMux := middlewares.RegisterDefault(mux, middlewares.CheckContentType)

	return wrappedMux
}
