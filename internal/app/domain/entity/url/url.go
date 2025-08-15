// Пакет entity включает структуры с бизнес логикой приложения.
package entity

// URL представляет пару сокращённого и оригинального URL.
type URL struct {
	ShortURL    string
	OriginalURL string
}

// URLData содержит URL и информацию о том, удалён ли он.
type URLData struct {
	URL       string
	IsDeleted bool
}

// DeleteBatch содержит данные для пакетного удаления URL-адресов для конкретного пользователя.
type DeleteBatch struct {
	ListURL []string
	UserID  string
}

// Batch представляет сущность с уникальным ID и связью между оригинальным и сокращённым URL.
type Batch struct {
	ID          string
	OriginalURL string
	ShortURL    string
}

// ServiceStats содержит количество сокращенных ссылок и юзеров в сервисе.
type ServiceStats struct {
	URLCount  int
	UserCount int
}
