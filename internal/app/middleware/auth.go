package middleware

import (
	"context"
	"net/http"

	"url-shortener/internal/app/services/auth"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		claims := &auth.Claims{}

		cookie, err := req.Cookie(authCookieName)
		if err != nil {
			claims = auth.NewClaims()
			token, err := claims.CreateJWTToken()
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			http.SetCookie(res, &http.Cookie{
				Name:  authCookieName,
				Value: token,
				Path:  "/",
			})
		} else {
			err = claims.Parse(cookie.Value)
			if err != nil {
				http.Error(res, err.Error(), http.StatusUnauthorized)
				return
			}

			err = claims.Valid()
			if err != nil {
				http.Error(res, err.Error(), http.StatusUnauthorized)
				return
			}
		}

		ctx := context.WithValue(req.Context(), userIDKey, claims.UserID)
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}
