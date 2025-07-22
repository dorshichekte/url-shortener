package urlrepositorypostgres

import (
	entity "url-shortener/internal/app/domain/entity/url"
)

func (s *urlRepositoryPostgres) DeleteBatch(event entity.DeleteBatch) error {
	query := `
        UPDATE urls
        SET is_deleted = true
        WHERE short_url = ANY($1::text[]) AND user_id = $2
    `

	_, err := s.db.Exec(query, event.ListURL, event.UserID)
	if err != nil {
		return err
	}

	return nil
}
