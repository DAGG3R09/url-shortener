package repository

import (
	"sync"

	"github.com/ansel1/merry"

	"github.com/DAGG3R09/url-shortener/models"
	"github.com/DAGG3R09/url-shortener/service/errors"
)

var _ RepositoryIface = (*MapRepository)(nil)

type MapRepository struct {
	urlStore     *sync.Map
	originalURLs *sync.Map
}

func NewMapRepository() *MapRepository {
	m := MapRepository{
		urlStore:     new(sync.Map),
		originalURLs: new(sync.Map),
	}

	return &m
}

func (mr *MapRepository) Exists(key int64) (exists bool) {
	_, ok := mr.urlStore.Load(key)

	return ok
}

func (mr *MapRepository) Store(urlBlock *models.URL) error {
	mr.urlStore.Store(urlBlock.Key, urlBlock)
	mr.originalURLs.Store(urlBlock.OriginalURL, urlBlock.Key)

	return nil
}

func (mr *MapRepository) Fetch(key int64) (url *models.URL, err error) {
	v, ok := mr.urlStore.Load(key)
	if !ok {
		return nil, merry.New("not found")
	}

	return v.(*models.URL), nil
}

func (mr *MapRepository) FetchKey(originalURL string) (key int64, err error) {
	v, ok := mr.originalURLs.Load(originalURL)
	if !ok {
		return 0, errors.SetError(errors.URLNotFoundError)
	}

	return v.(int64), nil
}

func (mr *MapRepository) FetchByURL(originalURL string) (url *models.URL, err error) {
	v, ok := mr.originalURLs.Load(originalURL)
	if !ok {
		return nil, errors.SetError(errors.URLNotFoundError)
	}

	v, ok = mr.urlStore.Load(v.(int64))
	if !ok {
		return nil, errors.SetError(errors.URLNotFoundError)
	}

	return v.(*models.URL), nil
}
