package dbusecase

import (
	"context"
)

// Ping проверка соединения с бд.
func (u *DBUseCase) Ping(ctx context.Context) error {
	err := u.dbConn.PingContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
