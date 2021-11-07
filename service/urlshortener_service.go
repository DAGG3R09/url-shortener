package service

import (
	"net/url"

	"github.com/labstack/gommon/log"

	"github.com/DAGG3R09/url-shortener/models"
	"github.com/DAGG3R09/url-shortener/repository"
	"github.com/DAGG3R09/url-shortener/service/errors"
	"github.com/DAGG3R09/url-shortener/service/keygenerator"
	"github.com/DAGG3R09/url-shortener/service/shorteners"
)

// var log = flume.New("urlshortener")

// URLShortenerService
// Goal: to shorten a given URL. Duplicate URLs should return the same shortened URL.
type URLShortnerService struct {
	shortener  shorteners.ShortenerIface
	repository repository.RepositoryIface

	generator *keygenerator.RandomNumber
}

func NewURLShortnerService(s shorteners.ShortenerIface,
	r repository.RepositoryIface, minRandomNumber, maxRandomNumber int64) *URLShortnerService {
	return &URLShortnerService{
		shortener:  s,
		repository: r,
		generator:  keygenerator.NewRandomNumber(minRandomNumber, maxRandomNumber),
	}
}

func (u *URLShortnerService) Shorten(originalURL string) (string, error) {
	log.Debug("in Shorten of URLShortnerService")

	if originalURL == "" {
		return "", errors.SetError(errors.URLNotProvidedError)
	}

	_, err := url.ParseRequestURI(originalURL)
	if err != nil {
		return "", errors.SetError(errors.InvalidURLError, "url", originalURL)
	}

	var (
		UniqueKeyNotFound bool = true
		key               int64
		shortURL          string
		url               *models.URL
	)

	// reuse the same short URL for the same URL
	url, err = u.repository.FetchByURL(originalURL)
	if err == nil {
		return url.ShortURL, nil
	}

	// we don't want keys that are already used in our repository
	for UniqueKeyNotFound {
		key = u.generator.Generate()

		UniqueKeyNotFound = u.repository.Exists(key)
	}

	log.Info("Key generated", "key", key, "url", originalURL)

	shortURL = u.shortener.Shorten(key)

	url = models.NewURL(originalURL, shortURL, key)

	err = u.repository.Store(url)
	if err != nil {
		return "", errors.SetError(errors.InternalServerError,
			"error", "failed to save to repository")
	}

	return shortURL, nil
}

func (u *URLShortnerService) FetchRedirectURL(shortURL string) (string, error) {
	key, err := u.shortener.Unshorten(shortURL)
	if err != nil {
		return "", errors.SetError(errors.InvalidURLError, "shortURL", shortURL)
	}

	urlObject, err := u.repository.Fetch(key)
	if err != nil {
		return "", errors.SetError(errors.URLNotFoundError, "shortURL", shortURL)
	}

	return urlObject.OriginalURL, nil
}

const (
	// We are generating base62 string of len 6
	// the smallest valued Base62 string would be "100000" and
	// the larged valued Base62 string would be "ZZZZZZ"
	// we have to make sure we generate a random number between this range.

	// min = Base10(Base62("100000")) = 916132832
	MinValueBase62Number int64 = 916132832

	// min = Base10(Base62("ZZZZZZ")) = 56800235583
	MaxValueBase62Number int64 = 56800235583
)
