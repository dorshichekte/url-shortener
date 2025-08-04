package graceful

import (
	"context"
	"log/slog"
)

// Process описывает единицу работы, которую можно запустить.
type Process struct {
	starter  starter
	disabled bool
}

type starter interface {
	Start(ctx context.Context) error
}

// Graceful управляет запуском группы процессов и логированием их состояния.
type Graceful struct {
	processes []Process
	logger    *slog.Logger
}
