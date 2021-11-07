package repository

import "github.com/DAGG3R09/url-shortener/models"

// type RepositoryIface interface {
// 	Store(originalURL string, key int64) error
// 	Fetch(string) (string, error)
// }

type RepositoryIfaceDeprecated interface {
	Store(originalURL string, key int64) error
	Fetch(key int64) (string, error)
	Exists(key int64) bool

	FetchKey(originalURL string) (key int64, err error)
}

type RepositoryIface interface {
	Store(url *models.URL) error
	Fetch(key int64) (*models.URL, error)
	Exists(key int64) bool

	// FetchKey(originalURL string) (key int64, err error)
	FetchByURL(originalURL string) (key *models.URL, err error)
}
