package handlers

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	urlS "url-shortener/internal/app/services/url"
)

func GetURL(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	path := req.URL.Path
	shortURL := strings.TrimPrefix(path, "/")

	originalURL, err := urlS.GetOriginal(shortURL)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.Header().Add("Location", originalURL)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func AddURL(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

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

	trimmedURL := strings.TrimSpace(string(body))
	_, err = url.ParseRequestURI(trimmedURL)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	shortURL, err := urlS.CreateShort(trimmedURL)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.Header().Add("Content-Type", "text/plain; charset=utf-8")
	res.WriteHeader(http.StatusCreated)
	_, err = res.Write([]byte(shortURL))

	if err != nil {
		return
	}
}
