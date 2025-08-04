package urlrepositorypostgres

import (
	"context"
	"database/sql"
	"errors"

	"url-shortener/internal/app/repository/model"
	"url-shortener/internal/pkg/constants"
)

// GetOriginalByID возвращает данные оригинального URL по короткому URL из базы данных.
func (s *urlRepositoryPostgres) GetOriginalByID(ctx context.Context, shortURL string) (model.URLData, error) {
	var URLData model.URLData

	query := `SELECT url, is_deleted
			  FROM urls
    		  WHERE short_url = $1
    `
	err := s.db.QueryRowContext(ctx, query, shortURL).Scan(&URLData.URL, &URLData.IsDeleted)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return URLData, constants.ErrURLNotFound
		}
		return URLData, err
	}
	return URLData, nil
}
