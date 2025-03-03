package handlers

import (
	"fmt"
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
	fmt.Println(shortURL)
	originalURL, err := urlS.GetOriginal(shortURL)
	fmt.Println(originalURL)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.Header().Set("Location", originalURL)
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

	parsedURL, err := url.Parse(string(body))
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	shortURL := urlS.CreateShort(string(body))
	baseURL := fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host)

	fullURL := baseURL + "/" + shortURL

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(fullURL))
}
