package dbusecase

import (
	"context"
	"database/sql"
)

// IDBUseCase описывает интерфейс бизнес-логики для работы с URL.
type IDBUseCase interface {
	Ping(ctx context.Context) error
}

type DBUseCase struct {
	dbConn *sql.DB
}
