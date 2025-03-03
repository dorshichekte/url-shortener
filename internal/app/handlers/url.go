package handlers

import (
	"io"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"url-shortener/internal/app/config"
	urlService "url-shortener/internal/app/services/url"
)

func Get(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	originalURL, err := urlService.GetOriginal(id)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.Header().Set("Location", originalURL)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func Add(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	if len(body) == 0 {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = url.ParseRequestURI(string(body))
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	shortURL := urlService.CreateShort(string(body))
	baseURL := config.GetConfig().BaseURL
	fullURL := baseURL + "/" + shortURL

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(fullURL))
}
