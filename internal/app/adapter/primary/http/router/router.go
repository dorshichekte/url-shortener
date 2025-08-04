// Пакет router предоставляет инструменты для маршрутизации HTTP-запросов.
package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	m "github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	"url-shortener/internal/app/adapter/primary/http/middleware"
)

// New создаёт новый экземпляр Router с заданными зависимостями.
func New(logger *zap.Logger) *Router {
	router := chi.NewRouter()

	r := &Router{
		router: router,
		logger: logger,
	}

	return r
}

// Router возвращает хэндлер
func (r *Router) Router() http.Handler {
	return r.router
}

func (r *Router) appendRoutesToRouter(subRouter *chi.Mux, routes []Route) {
	globalMiddlewares := chi.Middlewares{middleware.Log(r.logger), middleware.Gzip, middleware.Decompress}

	subRouter.Use(globalMiddlewares...)
	subRouter.Mount("/profiler", m.Profiler())

	for _, route := range routes {
		subRouter.Method(route.Method, route.Path, route.Handler)
	}
}
