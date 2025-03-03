package handlers

import (
	"net/http"

	"url-shortener/internal/app/handlers/middlewares"
)

func Register() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", AddUrl)
	mux.HandleFunc("/{id}", GetUrl)

	wrappedMux := middlewares.RegisterDefault(mux, middlewares.CheckContentType)

	return wrappedMux
}
