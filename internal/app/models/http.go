package models

type ShortenRequest struct {
	OriginalURL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"result"`
}

type BatchRequest struct {
	Id          string `json:"correlation_id"`
	OriginalURL string `json:"original_url"`
}

type BatchResponse struct {
	Id       string `json:"correlation_id"`
	ShortURL string `json:"short_url"`
}
