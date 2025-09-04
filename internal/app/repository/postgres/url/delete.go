package urlrepositorypostgres

import "context"

// Delete удаляет запись из таблицы urls по короткому URL.
func (u *urlRepositoryPostgres) Delete(ctx context.Context, shortURL string) error {
	query := `DELETE FROM urls
       		  WHERE short_url = $1
       		  AND url = $1
    `

	_, err := u.db.ExecContext(ctx, query, shortURL)
	if err != nil {
		return err
	}

	return nil
}
