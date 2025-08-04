package middleware

import (
	"io"
	"net/http"
)

// Middleware тип-алиас для функций middleware.
type Middleware = func(f http.Handler) http.Handler

type responseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

type gzipWriter struct {
	io.Writer
	http.ResponseWriter
}

// ContextKey тип для ключей контекста.
type ContextKey string
