package app

import httpadapter "url-shortener/internal/app/adapter/primary/http"

// App представляет основное приложение с HTTP-адаптером.
type App struct {
	HTTPAdapter *httpadapter.HTTPAdapter
}
