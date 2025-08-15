package urlrepositorypostgres

import (
	"context"
	"fmt"

	"url-shortener/internal/app/repository/model"
)

// GetAllByUserID возвращает список всех URL, принадлежащих пользователю с userID.
func (u *urlRepositoryPostgres) GetAllByUserID(ctx context.Context, userID string) ([]model.URL, error) {
	var listURLs []model.URL

	query := `SELECT url, short_url 
   			  FROM urls 
    		  WHERE user_id = $1
    `
	rows, err := u.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var url model.URL
		if err = rows.Scan(&url.OriginalURL, &url.ShortURL); err != nil {
			return nil, err
		}
		url.ShortURL = fmt.Sprintf("%s/%s", u.config.BaseURL, url.ShortURL)
		listURLs = append(listURLs, url)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return listURLs, nil
}
