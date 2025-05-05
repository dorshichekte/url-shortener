package service

import (
	"url-shortener/internal/app/service/url"
	"url-shortener/internal/app/service/worker"
)

type Services struct {
	URL    url.Methods
	Worker *worker.Service
}
