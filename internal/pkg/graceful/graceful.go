// Пакет graceful завершает процессы в приложении.
package graceful

import (
	"context"
	"log/slog"

	"golang.org/x/sync/errgroup"
)

// New создает экземпляр Graceful с переданными процессами и стандартным логгером.
func New(processes ...Process) *Graceful {
	return &Graceful{
		processes: processes,
		logger:    slog.Default(),
	}
}

// Start запускает все процессы параллельно с помощью errgroup.
func (gr *Graceful) Start(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	for _, process := range gr.processes {
		process := process

		if process.disabled {
			continue
		}

		f := func() error {
			err := process.starter.Start(ctx)
			if err != nil {
				gr.logger.Error(err.Error())
				return err
			}

			return nil
		}

		g.Go(f)
	}

	err := g.Wait()
	if err != nil {
		return err
	}

	return nil
}

// SetLogger позволяет установить пользовательский логгер вместо slog.Default().
func (gr *Graceful) SetLogger(l *slog.Logger) {
	gr.logger = l
}
