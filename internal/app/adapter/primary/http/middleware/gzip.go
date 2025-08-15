package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	errorhandler "url-shortener/internal/app/adapter/primary/http/handler/errors"
)

// Gzip сжимает ответ в gzip, если клиент поддерживает.
func Gzip(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		acceptEncoding := r.Header.Get("Accept-Encoding")
		isGzip := strings.Contains(acceptEncoding, "gzip")

		if isGzip {
			contentType := r.Header.Get("Content-Type")
			isAcceptContentType := contentType == "application/json" || contentType == "text/html"

			if isAcceptContentType {
				w.Header().Set("Content-Encoding", "gzip")

				gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
				if err != nil {
					_, _ = io.WriteString(w, err.Error())
					return
				}
				defer func() {
					_ = gz.Close()
				}()

				gzWriter := &gzipWriter{Writer: gz, ResponseWriter: w}
				w = gzWriter
			}
		}

		next.ServeHTTP(w, r)
	})
}

// Write записывает данные.
func (w gzipWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// Decompress распаковывает gzip-запросы.
func Decompress(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentEncoding := r.Header.Get("Content-Encoding")
		isGzipEncoding := contentEncoding == "gzip"

		if isGzipEncoding {
			zr, err := gzip.NewReader(r.Body)
			if err != nil {
				http.Error(w, errorhandler.ErrMessageFailedDecompressBody, http.StatusBadRequest)
				return
			}
			defer func() {
				_ = zr.Close()
			}()

			r.Body = http.MaxBytesReader(w, zr, 10<<20)
		}

		next.ServeHTTP(w, r)
	})
}
