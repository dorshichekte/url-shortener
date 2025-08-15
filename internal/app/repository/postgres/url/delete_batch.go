package urlrepositorypostgres

import (
	entity "url-shortener/internal/app/domain/entity/url"
)

// DeleteBatch помечает как удалённые (is_deleted = true) записи в таблице urls,
func (u *urlRepositoryPostgres) DeleteBatch(event entity.DeleteBatch) error {
	query := `
        UPDATE urls
        SET is_deleted = true
        WHERE short_url = ANY($1::text[]) AND user_id = $2
    `

	_, err := u.db.Exec(query, event.ListURL, event.UserID)
	if err != nil {
		return err
	}

	return nil
}
