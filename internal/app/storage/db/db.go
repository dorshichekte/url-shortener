package db

import (
	"database/sql"
	"errors"

	_ "github.com/jackc/pgx/v5/stdlib"

	"url-shortener/internal/app/constants"
)

func NewPostgresStorage(dsn string) (*Storage, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	createTableQuery := `
    CREATE TABLE IF NOT EXISTS urls (
        id SERIAL PRIMARY KEY,
        url TEXT NOT NULL,
        short_url TEXT NOT NULL UNIQUE
    );
    `
	if _, err = db.Exec(createTableQuery); err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

func (r *Storage) Get(shortURL string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var url string
	err := r.db.QueryRow("SELECT url FROM urls WHERE short_url = $1", shortURL).Scan(&url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", constants.ErrURLNotFound
		}
		return "", err
	}

	return url, nil
}

func (r *Storage) Add(url string, shortURL string) {
	query := "INSERT INTO urls (url, short_url) VALUES ($1, $2)"
	r.db.Exec(query, shortURL, url)
}

func (r *Storage) Delete(shortURL string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	query := "DELETE FROM urls WHERE short_url = $1 AND url = $1"
	_, err := r.db.Exec(query, shortURL)

	if err != nil {
		return err
	}

	return nil
}
