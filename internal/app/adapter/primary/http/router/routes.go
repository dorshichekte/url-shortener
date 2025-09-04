package router

import (
	"net/http"

	"url-shortener/internal/app/adapter/primary/http/handler"
	"url-shortener/internal/app/adapter/primary/http/middleware"
	"url-shortener/internal/app/config"
	"url-shortener/internal/pkg/auth"
)

// AppendRoutes добавляет пути приложения
func (r *Router) AppendRoutes(config *config.Config, handlers *handler.Handlers, auth auth.Auth) {
	r.config = &config.HTTPAdapter.Router

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
		{
			Path:    "/api/internal/stats",
			Method:  http.MethodGet,
			Handler: middleware.Add(middleware.Auth(auth), middleware.Subnet(config.Env))(http.HandlerFunc(handlers.URL.GetStats)),
		},
	}

	r.appendRoutesToRouter(r.router, routes)
}
