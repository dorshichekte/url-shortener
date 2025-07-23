// Пакет Dto включает структуры в которые преобразуются данные запроса и ответа в ручках.
package dto

// ShortenRequest представляет входные данные для запроса сокращения URL.
type ShortenRequest struct {
	OriginalURL string `json:"url"`
}

// ShortenResponse представляет ответ API после успешного сокращения URL.
type ShortenResponse struct {
	ShortURL string `json:"result"`
}

// BatchRequest представляет элемент запроса для пакетного сокращения URL.
type BatchRequest struct {
	ID          string `json:"correlation_id" validate:"required,min=1"`
	OriginalURL string `json:"original_url" validate:"required"`
}

// BatchResponse представляет элемент ответа при пакетном сокращении URL.
type BatchResponse struct {
	ID       string `json:"correlation_id"`
	ShortURL string `json:"short_url"`
}

// URLRequest используется для передачи пары "оригинальный URL — сокращённый URL".
type URLRequest struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

// DeleteBatchRequest представляет запрос на пометку URL как удалённых.
type DeleteBatchRequest struct {
	ListURL []string `json:"list_url"`
	UserID  string   `json:"user_id"`
}
