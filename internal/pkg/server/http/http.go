// Пакет server инициализирует хттп сервер.
package httpserver

import (
	"context"
	"errors"
	"net/http"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"url-shortener/internal/app/config"
)

// New создает и настраивает новый HTTP-сервер.
func New(logger *zap.Logger, config *config.Config, handler http.Handler) *Http {
	server := &http.Server{
		Handler:           handler,
		ReadTimeout:       config.HTTPAdapter.Server.ReadTimeout,
		WriteTimeout:      config.HTTPAdapter.Server.WriteTimeout,
		ReadHeaderTimeout: config.HTTPAdapter.Server.ReadHeaderTimeout,
		Addr:              config.HTTPAdapter.Server.Address,
	}

	s := Http{
		logger: logger,
		server: server,
		config: config,
	}

	return &s
}

// Start запускает HTTP-сервер и отслеживает завершение через контекст.
func (h *Http) Start(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), h.config.HTTPAdapter.Server.ShutdownTimeout)
		defer cancel()

		err := h.server.Shutdown(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	g.Go(func() error {
		var err error
		if h.config.Env.EnableHTTPS {
			err = h.server.ListenAndServeTLS("certs/cert.pem", "certs/key.pem")
		} else {
			err = h.server.ListenAndServe()
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
