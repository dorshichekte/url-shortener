package urlrepositorypostgres

import (
	"context"

	entity "url-shortener/internal/app/domain/entity/url"
)

// AddBatch выполняет пакетную вставку нескольких записей URL в базу данных Postgres.
func (u *urlRepositoryPostgres) AddBatch(ctx context.Context, batches []entity.Batch, userID string) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	query := `INSERT INTO urls (short_url, url, user_id) 
			  VALUES ($1, $2, $3)
    `
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer func() {
		_ = stmt.Close()
	}()

	for _, batch := range batches {
		_, err = stmt.ExecContext(ctx, batch.ShortURL, batch.OriginalURL, userID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
