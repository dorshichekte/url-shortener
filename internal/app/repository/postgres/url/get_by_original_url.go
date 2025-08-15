package urlrepositorypostgres

import (
	"context"
)

// GetByOriginalURL возвращает короткий URL по оригинальному URL из базы данных.
func (u *urlRepositoryPostgres) GetByOriginalURL(ctx context.Context, originalURL string) (string, error) {
	query := `SELECT short_url
			  FROM urls
			  WHERE url = $1
	`
	var shortURL string
	err := u.db.QueryRowContext(ctx, query, originalURL).Scan(&shortURL)
	if err != nil {
		return "", err
	}

	return shortURL, err
}
