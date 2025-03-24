package url

import (
	"url-shortener/internal/app/config"
	"url-shortener/internal/app/services/url"
)

type Handler struct {
	service *url.Service
	config  *config.Config
}
