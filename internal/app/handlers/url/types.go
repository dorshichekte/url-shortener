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
