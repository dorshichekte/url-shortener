package urlrepositorypostgres

import "context"

func (s *urlRepositoryPostgres) Delete(ctx context.Context, shortURL string) error {
	query := `DELETE FROM urls
       		  WHERE short_url = $1
       		  AND url = $1
    `

	_, err := s.db.ExecContext(ctx, query, shortURL)

	if err != nil {
		return err
	}

	return nil
}
