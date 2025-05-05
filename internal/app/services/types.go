package services

import (
	"url-shortener/internal/app/services/url"
	"url-shortener/internal/app/services/worker"
)

type Services struct {
	URL    url.Methods
	Worker *worker.Service
}
