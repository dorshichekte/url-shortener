package middleware

import (
	"net/http"

	config "url-shortener/internal/app/config/env"
)

// Subnet проверка доступа к методу по ip адресу
func Subnet(config *config.Env) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			if config.TrustedSubnet == "" {
				res.WriteHeader(http.StatusForbidden)
				return
			}

			ipHeader := req.Header.Get("X-Real-Ip")
			if ipHeader == "" {
				res.WriteHeader(http.StatusForbidden)
				return
			}

			if config.TrustedSubnet != ipHeader {
				res.WriteHeader(http.StatusForbidden)
				return
			}

			next.ServeHTTP(res, req)
		})
	}
}
