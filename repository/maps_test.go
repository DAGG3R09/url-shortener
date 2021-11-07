package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DAGG3R09/url-shortener/models"
)

func Test_mapsRepository(t *testing.T) {
	var m = NewMapRepository()
	assert.NotNil(t, m)

	url := "https://google.com"
	key := int64(907546489)
	shortURL := "123456"

	t.Run("store", func(t *testing.T) {
		err := m.Store(models.NewURL(url, shortURL, key))
		assert.NoError(t, err)

		fetchedURL, err := m.Fetch(key)
		assert.NotNil(t, fetchedURL)
		assert.NoError(t, err)
		assert.Equal(t, url, fetchedURL.OriginalURL)
		assert.Equal(t, shortURL, fetchedURL.ShortURL)
	})

	t.Run("Exist test", func(t *testing.T) {
		e := m.Exists(key)
		assert.True(t, e)

		e = m.Exists(11111111)
		assert.False(t, e)
	})

	t.Run("fetch Test", func(t *testing.T) {
		fetchedURL, err := m.Fetch(11111111)
		assert.Error(t, err)
		assert.Nil(t, fetchedURL)

		fetchedURL, err = m.Fetch(key)
		assert.NoError(t, err)
		assert.Equal(t, key, fetchedURL.Key)
		assert.Equal(t, url, fetchedURL.OriginalURL)
		assert.Equal(t, shortURL, fetchedURL.ShortURL)
	})

	t.Run("fetchKey Test", func(t *testing.T) {
		fetchedKey, err := m.FetchKey("invalid-URL")
		assert.Error(t, err)
		assert.Empty(t, fetchedKey)

		fetchedKey, err = m.FetchKey(url)
		assert.NoError(t, err)
		assert.Equal(t, key, fetchedKey)
	})

	t.Run("fetchByURL", func(t *testing.T) {
		urlBlock, err := m.FetchByURL("invalid-URL")
		assert.Error(t, err)
		assert.Nil(t, urlBlock)

		urlBlock, err = m.FetchByURL(url)
		assert.NoError(t, err)
		assert.Equal(t, key, urlBlock.Key)
		assert.Equal(t, url, urlBlock.OriginalURL)
		assert.Equal(t, shortURL, urlBlock.ShortURL)
	})

}
