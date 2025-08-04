// Пакет model содержит модели для таблиц базы данных.
package model

// Batch представляет запись для пакетной вставки URL-ов.
type Batch struct {
	ID          string `db:"id"`
	OriginalURL string `db:"original_url"`
	ShortURL    string `db:"short_url"`
}

// Shorter содержит данные для сокращённого URL,
type Shorter struct {
	OriginalURL string `db:"original_url"`
	ShortURL    string `db:"short_url"`
	UserID      string `db:"user_id"`
	Deleted     bool   `db:"is_deleted"`
}

// URLData хранит URL и информацию о его удалении.
type URLData struct {
	URL       string `db:"url"`
	IsDeleted bool   `db:"is_is_deleted"`
}

// URL описывает пару оригинального и короткого URL,
type URL struct {
	OriginalURL string `db:"original_url"`
	ShortURL    string `db:"short_url"`
}

// DeleteEvent описывает событие удаления URL-ов,
type DeleteEvent struct {
	ListURL []string `db:"list_url"`
	UserID  string   `db:"user_id"`
}
