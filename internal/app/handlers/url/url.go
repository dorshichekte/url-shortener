package url

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"url-shortener/internal/app/config"
	"url-shortener/internal/app/constants"
	urlService "url-shortener/internal/app/services/url"
)

func NewHandler(service *urlService.Service, cfg *config.Config) *Handler {
	return &Handler{
		service: service,
		config:  cfg,
	}
}

func (h *Handler) handleError(res http.ResponseWriter, statusCode int) {
	res.WriteHeader(statusCode)
}

func (h *Handler) jsonDecode(req *http.Request) (ShortenRequest, error) {
	var request ShortenRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		return request, err
	}

	return request, nil
}

func (h *Handler) parseRequestURL(req *http.Request) (string, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return "", err
	}
	defer req.Body.Close()

	if len(body) == 0 {
		return "", constants.ErrEmptyRequestBody
	}

	_, err = url.ParseRequestURI(string(body))
	if err != nil {
		return "", err
	}
	fmt.Println(body)

	return string(body), nil
}

func (h *Handler) Get(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	originalURL, err := h.service.GetOriginal(id)
	if err != nil {
		h.handleError(res, http.StatusBadRequest)
		return
	}

	res.Header().Set("Location", originalURL)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *Handler) Add(res http.ResponseWriter, req *http.Request) {
	originalURL, err := h.parseRequestURL(req)
	if err != nil {
		h.handleError(res, http.StatusBadRequest)
		return
	}

	shortURL := h.service.CreateShort(originalURL)
	baseURL := h.config.BaseURL
	fullURL := baseURL + "/" + shortURL

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(fullURL))
}

func (h *Handler) Shorten(res http.ResponseWriter, req *http.Request) {
	url, err := h.jsonDecode(req)
	if err != nil {
		h.handleError(res, http.StatusBadRequest)
		return
	}

	shortURL := h.service.CreateShort(url.OriginalURL)
	baseURL := h.config.BaseURL
	fullURL := baseURL + "/" + shortURL

	response := ShortenResponse{
		ShortURL: fullURL,
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	if resErr := json.NewEncoder(res).Encode(response); resErr != nil {
		h.handleError(res, http.StatusInternalServerError)
		return
	}
}
