// Package Dto contains a structure that transforms data for sending to another layer.
package dto

type ShortenRequest struct {
	OriginalURL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"result"`
}

type BatchRequest struct {
	ID          string `json:"correlation_id" validate:"required,min=1"`
	OriginalURL string `json:"original_url" validate:"required"`
}

type BatchResponse struct {
	ID       string `json:"correlation_id"`
	ShortURL string `json:"short_url"`
}

type URLRequest struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type DeleteBatchRequest struct {
	ListURL []string `json:"list_url"`
	UserID  string   `json:"user_id"`
}
