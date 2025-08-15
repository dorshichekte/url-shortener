// Пакет server инициализирует хттп сервер.
package server

import (
	"context"
	"errors"
	"net/http"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"url-shortener/internal/app/config"
)

// New создает и настраивает новый HTTP-сервер.
func New(logger *zap.Logger, config *config.Config, handler http.Handler) *Server {
	server := &http.Server{
		Handler:           handler,
		ReadTimeout:       config.HTTPAdapter.Server.ReadTimeout,
		WriteTimeout:      config.HTTPAdapter.Server.WriteTimeout,
		ReadHeaderTimeout: config.HTTPAdapter.Server.ReadHeaderTimeout,
		Addr:              config.HTTPAdapter.Server.Address,
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

		ctx, cancel := context.WithTimeout(context.Background(), a.config.HTTPAdapter.Server.ShutdownTimeout)
		defer cancel()

		err := a.server.Shutdown(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	g.Go(func() error {
		var err error
		if a.config.Env.EnableHTTPS {
			err = a.server.ListenAndServeTLS("certs/cert.pem", "certs/key.pem")
		} else {
			err = a.server.ListenAndServe()
		}
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
