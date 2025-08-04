package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

// Log middleware для логирования HTTP-запросов и ответов.
func Log(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			logger.Info("Request",
				zap.String("method", r.Method),
				zap.String("uri", r.RequestURI),
			)

			rw := &responseWriter{ResponseWriter: w}
			next.ServeHTTP(rw, r)

			logger.Info("Response",
				zap.Int("status", rw.status),
				zap.Int("size", rw.size),
				zap.Duration("duration", time.Since(start)),
			)
		})
	}
}

// WriteHeader перехватывает статус ответа для логирования.
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

// Write перехватывает размер данных для логирования.
func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += size
	return size, err
}
