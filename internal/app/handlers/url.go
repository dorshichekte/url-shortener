package handlers

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	urlS "url-shortener/internal/app/services/url"
)

func GetUrl(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	path := req.URL.Path
	shortUrl := strings.TrimPrefix(path, "/")

	originalUrl, err := urlS.GetOriginal(shortUrl)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.Header().Add("Location", originalUrl)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func AddUrl(res http.ResponseWriter, req *http.Request) {
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

	trimmedUrl := strings.TrimSpace(string(body))
	_, err = url.ParseRequestURI(trimmedUrl)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	shortUrl, err := urlS.CreateShort(trimmedUrl)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.Header().Add("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	_, err = res.Write([]byte(shortUrl))

	if err != nil {
		return
	}
}
