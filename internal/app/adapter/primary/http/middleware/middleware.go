package middleware

import "net/http"

func UserIDKey() ContextKey {
	return userIDKey
}

func Add(middlewares ...Middleware) Middleware {
	return func(h http.Handler) http.Handler {
		wrapped := h

		for _, middleware := range middlewares {
			wrapped = middleware(wrapped)
		}

		return wrapped
	}
}
