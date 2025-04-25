package db

import (
	"database/sql"
	"sync"
	"url-shortener/internal/app/config"
)

type Storage struct {
	db  *sql.DB
	cfg config.Config
	mu  sync.RWMutex
}
