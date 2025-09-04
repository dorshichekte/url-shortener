package urlrepositorypostgres

import "context"

// GetStats возвращает количество сокращенных урл, и количество уникальных юзеров.
func (u *urlRepositoryPostgres) GetStats(ctx context.Context) (int, int, error) {
	query := `SELECT 
    (SELECT COUNT(*) FROM urls),
    (SELECT COUNT(DISTINCT user_id) FROM urls)
    `

	var urlCount, userCount int
	err := u.db.QueryRowContext(ctx, query).Scan(&urlCount, &userCount)
	if err != nil {
		return 0, 0, err
	}

	return urlCount, userCount, err
}
