package db

import (
	"database/sql"
	"sync"
)

type Storage struct {
	db *sql.DB
	mu sync.RWMutex
}
