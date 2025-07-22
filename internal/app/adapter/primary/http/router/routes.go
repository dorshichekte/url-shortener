package router

import (
	"net/http"

	"url-shortener/internal/app/adapter/primary/http/handler"
	"url-shortener/internal/app/adapter/primary/http/middleware"
	adapter "url-shortener/internal/app/config/adapter"
	"url-shortener/internal/pkg/auth"
)

func (r *Router) AppendRoutes(config *adapter.Router, handlers *handler.Handlers, auth auth.Auth) {
	r.config = config

	routes := []Route{
		{
			Path:    "/ping",
			Method:  http.MethodGet,
			Handler: http.HandlerFunc(handlers.Database.Ping),
		},
		{
			Path:    "/",
			Method:  http.MethodPost,
			Handler: middleware.Add(middleware.Auth(auth))(http.HandlerFunc(handlers.URL.AddShorten)),
		},
		{
			Path:    "/{id}",
			Method:  http.MethodGet,
			Handler: middleware.Add(middleware.Auth(auth))(http.HandlerFunc(handlers.URL.GetOriginalByID)),
		},
		{
			Path:    "/api/user/urls",
			Method:  http.MethodGet,
			Handler: middleware.Add(middleware.Auth(auth))(http.HandlerFunc(handlers.URL.GetAllByUserID)),
		},
		{
			Path:    "/api/shorten/batch",
			Method:  http.MethodPost,
			Handler: middleware.Add(middleware.Auth(auth))(http.HandlerFunc(handlers.URL.AddBatch)),
		},
		{
			Path:    "/api/shorten",
			Method:  http.MethodPost,
			Handler: middleware.Add(middleware.Auth(auth))(http.HandlerFunc(handlers.URL.MakeFromJSON)),
		},
		{
			Path:    "/api/user/urls",
			Method:  http.MethodDelete,
			Handler: middleware.Add(middleware.Auth(auth))(http.HandlerFunc(handlers.URL.DeleteBatch)),
		},
	}

	r.appendRoutesToRouter(r.router, routes)
}
