package db

import (
	"database/sql"

	"url-shortener/internal/app/common"
)

type Storage struct {
	db *sql.DB
	common.BaseStorageDependency
}
