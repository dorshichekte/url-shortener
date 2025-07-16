package graceful

import (
	"context"
	"log/slog"
)

type Process struct {
	starter  starter
	disabled bool
}

type starter interface {
	Start(ctx context.Context) error
}

type Graceful struct {
	processes []Process
	logger    *slog.Logger
}
