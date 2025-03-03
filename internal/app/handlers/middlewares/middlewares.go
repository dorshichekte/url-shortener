package middlewares

import (
	"net/http"

	"url-shortener/internal/app/constants"
)

type Middleware func(http.Handler) http.Handler

func CheckContentType(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isContentTypeTextPlain := r.Header.Get("Content-type") == "text/plain"

		if !isContentTypeTextPlain {
			http.Error(w, constants.TextPlain, http.StatusUnsupportedMediaType)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func RegisterDefault(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, m := range middlewares {
		h = m(h)
	}

	return h
}
