// Пакет server инициализирует хттп сервер.
package server

import (
	"context"
	"errors"
	"net/http"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	adapter "url-shortener/internal/app/config/adapter"
)

// New создает и настраивает новый HTTP-сервер.
func New(logger *zap.Logger, config *adapter.HTTPServer, handler http.Handler) *Server {
	server := &http.Server{
		Handler:           handler,
		ReadTimeout:       config.ReadTimeout,
		WriteTimeout:      config.WriteTimeout,
		ReadHeaderTimeout: config.ReadHeaderTimeout,
		Addr:              config.Address,
	}

	s := Server{
		logger: logger,
		server: server,
		config: config,
	}

	return &s
}

// Start запускает HTTP-сервер и отслеживает завершение через контекст.
func (a *Server) Start(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), a.config.ShutdownTimeout)
		defer cancel()

		err := a.server.Shutdown(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	g.Go(func() error {
		err := a.server.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				// ok
			} else {
				return err
			}
		}

		return nil
	})

	err := g.Wait()
	if err != nil {
		return err
	}

	return nil
}
