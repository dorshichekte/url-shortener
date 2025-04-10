package url

import (
	"go.uber.org/zap"
	"url-shortener/internal/app/config"
	"url-shortener/internal/app/services/url"
)

type Handler struct {
	service *url.Service
	config  *config.Config
	logger  *zap.Logger
}

type ShortenRequest struct {
	OriginalURL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"result"`
}

type ShortenBatchRequest struct {
	ID          string `json:"correlation_id"`
	OriginalURL string `json:"original_url"`
}

type ShortenBatchResponse struct {
	ID       string `json:"correlation_id"`
	ShortURL string `json:"short_url"`
}
