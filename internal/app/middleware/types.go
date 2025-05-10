package middleware

import (
	"io"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

type gzipWriter struct {
	io.Writer
	http.ResponseWriter
}

type contextKey string
