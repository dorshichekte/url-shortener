package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"

	"url-shortener/internal/app/config"
	"url-shortener/internal/app/constants"
	"url-shortener/internal/app/models"
)

func NewPostgresStorage(cfg config.Config) (*Storage, error) {
	db, err := sql.Open("pgx", cfg.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	err = applyMigrations(cfg)
	if err != nil {
		return nil, err
	}

	return &Storage{db: db, cfg: cfg}, nil
}

func applyMigrations(cfg config.Config) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %v", err)
	}

	migrationDirPath := "file://" + filepath.Join(wd, "migrations")

	m, err := migrate.New(migrationDirPath, cfg.DatabaseDSN)
	if err != nil {
		return fmt.Errorf("failed to initialize migration: %v", err)
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		return fmt.Errorf("failed to apply migrations: %v", err)
	}

	return nil
}

func (s *Storage) Get(shortURL string) (models.URLData, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var URLData models.URLData
	err := s.db.QueryRow("SELECT url, is_deleted FROM urls WHERE short_url = $1", shortURL).Scan(&URLData.URL, &URLData.Deleted)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return URLData, constants.ErrURLNotFound
		}
		return URLData, err
	}
	return URLData, nil
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
		url.ShortURL = fmt.Sprintf("%s/%s", s.cfg.BaseURL, url.ShortURL)
		listURLs = append(listURLs, url)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return listURLs, nil
}

func (s *Storage) BatchUpdate(shortURLs []string, userID string) error {
	query := `
        UPDATE urls
        SET is_deleted = true
        WHERE short_url = ANY($1::text[]) AND user_id = $2
    `
	_, err := s.db.Exec(query, shortURLs, userID)

	return err
}
