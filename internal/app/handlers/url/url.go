package url

import (
	"io"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"url-shortener/internal/app/config"
	"url-shortener/internal/app/constants"
	u "url-shortener/internal/app/services/url"
)

func NewHandler(service *u.Service, cfg *config.Config) *Handler {
	return &Handler{
		service: service,
		config:  cfg,
	}
}
func (h *Handler) handleError(res http.ResponseWriter, statusCode int) {
	res.WriteHeader(statusCode)
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
