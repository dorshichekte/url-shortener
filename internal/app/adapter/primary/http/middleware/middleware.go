// Пакет middleware включает все мидлварины приложения.
package middleware

import "net/http"

// Add добавляет миддварины к запросу.
func Add(middlewares ...Middleware) Middleware {
	return func(h http.Handler) http.Handler {
		wrapped := h

		for _, middleware := range middlewares {
			wrapped = middleware(wrapped)
		}

		return wrapped
	}
}
