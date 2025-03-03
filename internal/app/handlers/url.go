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

	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.Write([]byte(originalURL))
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

	_, err = url.ParseRequestURI(string(body))
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	shortURL := urlS.CreateShort(string(body))

	var fullURL string

	if strings.HasSuffix(string(body), "/") {
		fullURL = string(body) + shortURL
	} else {
		fullURL = string(body) + "/" + shortURL
	}

	res.Header().Add("Content-Type", "text/plain; charset=utf-8")
	res.WriteHeader(http.StatusCreated)
	_, err = res.Write([]byte(fullURL))

	if err != nil {
		return
	}
}
