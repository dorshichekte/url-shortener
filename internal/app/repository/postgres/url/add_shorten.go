package urlrepositorypostgres

import (
	"context"
)

// ToDo переделать логику
func (s *urlRepositoryPostgres) AddShorten(ctx context.Context, originalURL, shortURL, userID string) (string, error) {
	query := `INSERT INTO urls (url, short_url, user_id) 
			  VALUES ($1, $2, $3)
    `
	_, err := s.db.ExecContext(ctx, query, originalURL, shortURL, userID)
	if err != nil {
		return "", err
	}

	return "", nil
}
