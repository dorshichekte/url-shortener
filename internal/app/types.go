package app

import httpadapter "url-shortener/internal/app/adapter/primary/http"

type App struct {
	HTTPAdapter *httpadapter.HTTPAdapter
}
