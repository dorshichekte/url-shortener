package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"url-shortener/internal/app/models"

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

func (s *Storage) Get(shortURL string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var url string
	err := s.db.QueryRow("SELECT url FROM urls WHERE short_url = $1", shortURL).Scan(&url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", constants.ErrURLNotFound
		}
		return "", err
	}

	return url, nil
}

func (s *Storage) Add(url string, shortURL string) {
	query := "INSERT INTO urls (url, short_url) VALUES ($1, $2)"
	s.db.Exec(query, shortURL, url)
}

func (s *Storage) Delete(shortURL string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	query := "DELETE FROM urls WHERE short_url = $1 AND url = $1"
	_, err := s.db.Exec(query, shortURL)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) AddBatch(listBatches []models.Batch) error {
	queries := make([]string, 0, len(listBatches))

	for _, batch := range listBatches {
		query := fmt.Sprintf("INSERT INTO urls (short_url, url) VALUES ('%s', '%s');", batch.ShortURL, batch.OriginalURL)
		queries = append(queries, query)
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	for _, query := range queries {
		_, err = tx.Exec(query)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
