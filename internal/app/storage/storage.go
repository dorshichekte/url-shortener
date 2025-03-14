package storage

func NewURLStorage() *URLStorage {
	return &URLStorage{
		mapURL:      make(map[string]string),
		mapShortURL: make(map[string]string),
	}
}

func (us *URLStorage) Get(url string, urlType URLType) string {
	switch urlType {
	case DefaultURLType:
		value, _ := us.mapURL[url]
		return value
	case ShortURLType:
		value, _ := us.mapShortURL[url]
		return value
	default:
		return ""
	}
}

func (us *URLStorage) Add(url, shortURL string) {
	us.mapURL[url] = shortURL
	us.mapShortURL[shortURL] = url
}
