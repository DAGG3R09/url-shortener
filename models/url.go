package models

type URL struct {
	OriginalURL string
	ShortURL    string
	Key         int64
}

func NewURL(OriginalURL, ShortURL string, key int64) *URL {
	return &URL{
		OriginalURL: OriginalURL,
		ShortURL:    ShortURL,
		Key:         key,
	}
}
