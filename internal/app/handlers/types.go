package handlers

import (
	"url-shortener/internal/app/handlers/url"
)

type Handler struct {
	urlHandler *url.Handler
}
