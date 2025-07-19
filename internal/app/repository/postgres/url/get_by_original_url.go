package urlrepositorypostgres

import (
	"context"
)

func (s *urlRepositoryPostgres) GetByOriginalURL(ctx context.Context, originalURL string) (string, error) {
	query := `SELECT short_url
			  FROM urls
			  WHERE url = $1
	`
	var shortURL string
	err := s.db.QueryRowContext(ctx, query, originalURL).Scan(&shortURL)
	if err != nil {
		return "", err
	}

	return shortURL, err
}
