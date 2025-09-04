package dbusecase

import "database/sql"

func New(conn *sql.DB) *DBUseCase {
	return &DBUseCase{
		dbConn: conn,
	}
}
