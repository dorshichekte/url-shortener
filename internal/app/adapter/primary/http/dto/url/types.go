// Пакет Dto включает структуры в которые преобразуются данные запроса и ответа в ручках.
package dto

// ShortenRequest входные данные для запроса сокращения URL.
type ShortenRequest struct {
	OriginalURL string `json:"url"`
}

// ShortenResponse ответ API после успешного сокращения URL.
type ShortenResponse struct {
	ShortURL string `json:"result"`
}

// BatchRequest входные данные для запроса пакетного сокращения URL.
type BatchRequest struct {
	ID          string `json:"correlation_id" validate:"required,min=1"`
	OriginalURL string `json:"original_url" validate:"required"`
}

// BatchResponse ответ API при пакетном сокращении URL.
type BatchResponse struct {
	ID       string `json:"correlation_id"`
	ShortURL string `json:"short_url"`
}

// URLRequest входные данные пары "оригинальный URL — сокращённый URL".
type URLRequest struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

// DeleteBatchRequest  входные данные для метода DeleteBatch.
type DeleteBatchRequest struct {
	ListURL []string `json:"list_url"`
	UserID  string   `json:"user_id"`
}

// ServiceStatsResponse ответ API для метода GetStats.
type ServiceStatsResponse struct {
	URLCount  int `json:"urls"`
	UserCount int `json:"users"`
}
