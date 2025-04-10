package url

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"

	"url-shortener/internal/app/config"
	"url-shortener/internal/app/constants"
	urlService "url-shortener/internal/app/services/url"
)

func NewHandler(service *urlService.Service, cfg *config.Config, logger *zap.Logger) *Handler {
	return &Handler{
		service: service,
		config:  cfg,
		logger:  logger,
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
		h.logger.Error("Failed get original URL", zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}

	res.Header().Set("Location", originalURL)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *Handler) Add(res http.ResponseWriter, req *http.Request) {
	originalURL, err := h.parseRequestURL(req)
	if err != nil {
		h.logger.Error("Failed parse request URL", zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}
	defer func() {
		_ = req.Body.Close()
	}()

	shortURL := h.service.CreateShort(originalURL, h.config)
	baseURL := h.config.BaseURL
	fullURL := baseURL + "/" + shortURL

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(fullURL))
}

func (h *Handler) Shorten(res http.ResponseWriter, req *http.Request) {
	u, err := h.jsonDecode(req)
	if err != nil {
		h.logger.Error("Failed decode json", zap.Error(err))
		h.handleError(res, http.StatusInternalServerError)
		return
	}

	shortURL := h.service.CreateShort(u.OriginalURL, h.config)
	baseURL := h.config.BaseURL
	fullURL := baseURL + "/" + shortURL

	response := ShortenResponse{
		ShortURL: fullURL,
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	if resErr := json.NewEncoder(res).Encode(response); resErr != nil {
		h.logger.Error("Failed encode json", zap.Error(resErr))
		h.handleError(res, http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Ping(res http.ResponseWriter, req *http.Request) {
	ps := h.config.DatabaseDSN

	db, err := sql.Open("pgx", ps)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer func() {
		_ = db.Close()
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
}

func (h *Handler) Batch(res http.ResponseWriter, req *http.Request) {
	var rq []ShortenBatchRequest
	var rs []ShortenBatchResponse

	dec := json.NewDecoder(req.Body)
	if err := dec.Decode(&rq); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer func() {
		_ = req.Body.Close()
	}()

	for _, batch := range rq {
		if batch.ID == "" || batch.OriginalURL == "" {
			continue
		}

		shortURL := h.service.CreateShort(batch.OriginalURL, h.config)
		baseURL := h.config.BaseURL
		fullURL := baseURL + "/" + shortURL

		resp := ShortenBatchResponse{
			ID:       batch.ID,
			ShortURL: fullURL,
		}

		rs = append(rs, resp)

	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	enc := json.NewEncoder(res)
	if err := enc.Encode(rs); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}
