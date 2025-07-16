package middleware

import (
	"context"
	"net/http"

	"url-shortener/internal/pkg/auth"
	"url-shortener/internal/pkg/util/error_response"
)

func Auth(auth auth.Auth) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			cookie, cookieErr := req.Cookie(authCookieName)
			if cookieErr != nil {
				token, err := auth.Generate(1)
				if err != nil {
					//ToDo добавить логирование
					return
				}
				ctx := context.WithValue(req.Context(), userIDKey, token.AccessToken)
				next.ServeHTTP(res, req.WithContext(ctx))
				return
			}

			userData, parseAccessTokenErr := auth.ParseAccessToken(cookie.Value)
			if parseAccessTokenErr != nil {
				util.WriteErrorResponse(res, http.StatusUnauthorized, util.WrapperError[string]{CustomError: parseAccessTokenErr.Error()})
				return
			}

			ctx := context.WithValue(req.Context(), userIDKey, userData.ID)
			next.ServeHTTP(res, req.WithContext(ctx))
		})
	}
}
