package handlers

import (
	"url-shortener/internal/app/handlers/db"
	"url-shortener/internal/app/handlers/url"
)

type Handlers struct {
	URL      url.Method
	Database db.Method
}

type Handler struct {
	Handlers *Handlers
}
