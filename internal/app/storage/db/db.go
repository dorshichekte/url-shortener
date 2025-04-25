package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"

	"url-shortener/internal/app/constants"
	"url-shortener/internal/app/models"
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
        url TEXT NOT NULL UNIQUE,
        short_url TEXT NOT NULL UNIQUE,
        user_id VARCHAR(255) NOT NULL
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

func (s *Storage) Add(url, shortURL, userID string) (string, error) {
	query := "INSERT INTO urls (url, short_url, user_id) VALUES ($1, $2, $3)"
	_, err := s.db.Exec(query, url, shortURL, userID)
	if err != nil {
		var shortURL string
		s.db.QueryRow("SELECT short_url FROM urls WHERE url = $1", url).Scan(&shortURL)
		return shortURL, err
	}

	return "", nil
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

func (s *Storage) AddBatch(listBatches []models.Batch, userID string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	stmt, err := tx.Prepare("INSERT INTO urls (short_url, url, user_id) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer func() {
		_ = stmt.Close()
	}()

	for _, batch := range listBatches {
		_, err = stmt.Exec(batch.ShortURL, batch.OriginalURL, userID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s *Storage) GetUsersURLsByID(userID string) ([]models.URL, error) {
	var listURLs []models.URL

	rows, err := s.db.Query(`SELECT url, short_url FROM urls WHERE user_id=$1`, userID)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var url models.URL
		if err = rows.Scan(&url.OriginalURL, &url.ShortURL); err != nil {
			return nil, err
		}
		url.ShortURL = fmt.Sprintf("%s/%s", "ss", url.ShortURL)
		listURLs = append(listURLs, url)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return listURLs, nil
}
