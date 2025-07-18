package model

type Batch struct {
	ID          string `db:"id"`
	OriginalURL string `db:"original_url"`
	ShortURL    string `db:"short_url"`
}

type Shorter struct {
	OriginalURL string `db:"original_url"`
	ShortURL    string `db:"short_url"`
	UserID      string `db:"user_id"`
	Deleted     bool   `db:"is_deleted"`
}

type URLData struct {
	URL       string `db:"url"`
	IsDeleted bool   `db:"is_is_deleted"`
}

type URL struct {
	OriginalURL string `db:"original_url"`
	ShortURL    string `db:"short_url"`
}

type DeleteEvent struct {
	ListURL []string `db:"list_url"`
	UserID  string   `db:"user_id"`
}
