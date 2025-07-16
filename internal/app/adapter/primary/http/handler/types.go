package handler

import (
	db_hanler "url-shortener/internal/app/adapter/primary/http/handler/db"
	url_hanler "url-shortener/internal/app/adapter/primary/http/handler/url"
)

type Handlers struct {
	URL      *url_hanler.Handler
	Database *db_hanler.Handler
}
