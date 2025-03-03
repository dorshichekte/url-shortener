package handlers

import (
	"net/http"
)

func Register() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", AddURL)
	mux.HandleFunc("/{id}", GetURL)

	return mux
}
