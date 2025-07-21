package urlrepositorypostgres

import (
	"context"
	"fmt"

	entity "url-shortener/internal/app/domain/entity/url"
)

func (s *urlRepositoryPostgres) DeleteBatch(ctx context.Context, event entity.DeleteBatch) error {
	query := `
        UPDATE urls
        SET is_deleted = true
        WHERE short_url = ANY($1::text[]) AND user_id = $2
    `

	fmt.Println(event)
	res, err := s.db.ExecContext(ctx, query, event.ListURL, event.UserID)
	fmt.Println(res)
	return err
}
