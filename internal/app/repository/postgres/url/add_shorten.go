package urlrepositorypostgres

import (
	"context"
)

func (s *urlRepositoryPostgres) AddShorten(ctx context.Context, originalURL, shortURL, userID string) (string, error) {
	query := `INSERT INTO urls (url, short_url, user_id) 
			  VALUES ($1, $2, $3)
    `
	_, err := s.db.ExecContext(ctx, query, originalURL, shortURL, userID)
	if err != nil {
		var shortURL string
		query := `SELECT short_url 
				  FROM urls 
				  WHERE url = $1;
	    `
		err := s.db.QueryRowContext(ctx, query, originalURL).Scan(&shortURL)
		if err != nil {
			return "", err
		}
		return shortURL, err
	}

	return "", nil
}
