package entity

type URL struct {
	ShortURL    string
	OriginalURL string
}

type URLData struct {
	URL       string
	IsDeleted bool
}

type DeleteBatch struct {
	ListURL []string
	UserID  string
}

type Batch struct {
	ID          string
	OriginalURL string
	ShortURL    string
}
