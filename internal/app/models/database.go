package models

type Batch struct {
	ID          string
	OriginalURL string
	ShortURL    string
}

type Shorter struct {
	OriginalURL string
	ShortURL    string
	UserID      string
	Deleted     bool
}
