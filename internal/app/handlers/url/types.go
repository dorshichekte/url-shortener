package url

import (
	"url-shortener/internal/app/config"
	"url-shortener/internal/app/services/url"
)

type Handler struct {
	service *url.Service
	config  *config.Config
}

type ShortenRequest struct {
	OriginalURL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"result"`
}
