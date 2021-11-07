package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DAGG3R09/url-shortener/repository"
	"github.com/DAGG3R09/url-shortener/service/shorteners"
)

func Test_URLShortener_Shortener(t *testing.T) {

	u := NewURLShortnerService(
		shorteners.NewBase62Shortener(),
		repository.NewMapRepository(),
		MinValueBase62Number, MaxValueBase62Number)

	t.Run("Shorten", func(t *testing.T) {
		t.Run("empty url", func(t *testing.T) {
			url, err := u.Shorten("")
			assert.Error(t, err)
			assert.Empty(t, url)
		})

		t.Run("malformed URL", func(t *testing.T) {
			url, err := u.Shorten("9masi, bshyhas djkppp")
			assert.Error(t, err)
			assert.Empty(t, url)
		})

		t.Run("normal URL", func(t *testing.T) {
			url, err := u.Shorten("https://google.com")
			assert.NoError(t, err)
			assert.NotEmpty(t, url)
		})

		t.Run("Return same URL", func(t *testing.T) {
			url, err := u.Shorten("https://google.com")
			assert.NoError(t, err)
			assert.NotEmpty(t, url)

			url2, err := u.Shorten("https://google.com")
			assert.NoError(t, err)
			assert.NotEmpty(t, url2)
			assert.Equal(t, url, url2)
		})

	})
}

func Test_FetchRedirectURL(t *testing.T) {
	u := NewURLShortnerService(
		shorteners.NewBase62Shortener(),
		repository.NewMapRepository(),
		MinValueBase62Number, MaxValueBase62Number)

	t.Run("not found", func(t *testing.T) {
		url, err := u.FetchRedirectURL("asdasdads")
		assert.Error(t, err)
		assert.Empty(t, url)
	})

	t.Run("Found", func(t *testing.T) {
		url := "https://google.com"

		shortURL, err := u.Shorten(url)
		assert.NoError(t, err)
		assert.NotEmpty(t, shortURL)

		redirectURL, err := u.FetchRedirectURL(shortURL)
		assert.NoError(t, err)
		assert.Equal(t, url, redirectURL)
	})

}
