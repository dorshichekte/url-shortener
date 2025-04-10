package storage

type URLStorage interface {
	Get(url string) (string, error)
	Add(url, shortURL string)
	Delete(url string) error
}
