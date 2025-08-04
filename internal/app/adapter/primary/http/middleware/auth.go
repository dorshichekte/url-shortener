package middleware

import (
	"context"
	"net/http"

	"url-shortener/internal/pkg/auth"
	util "url-shortener/internal/pkg/util/error_response"
	stringUtils "url-shortener/internal/pkg/util/string"
)

func Auth(auth auth.Auth) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			cookie, cookieErr := req.Cookie(authCookieName)
			if cookieErr != nil {
				id := stringUtils.CreateRandom()

				token, err := auth.Generate(id)
				if err != nil {
					util.WriteErrorResponse(res, http.StatusUnauthorized, util.WrapperError[string]{CustomError: err.Error()})
					return
				}

				http.SetCookie(res, &http.Cookie{
					Name:  authCookieName,
					Value: token.AccessToken,
					Path:  "/",
				})

				ctx := context.WithValue(req.Context(), UserIDKey, id)
				next.ServeHTTP(res, req.WithContext(ctx))
				return
			}

			userData, parseAccessTokenErr := auth.ParseAccessToken(cookie.Value)
			if parseAccessTokenErr != nil {
				util.WriteErrorResponse(res, http.StatusUnauthorized, util.WrapperError[string]{CustomError: parseAccessTokenErr.Error()})
				return
			}

			ctx := context.WithValue(req.Context(), UserIDKey, userData.ID)
			next.ServeHTTP(res, req.WithContext(ctx))
		})
	}
}
