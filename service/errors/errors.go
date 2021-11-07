package errors

import (
	"fmt"
	"net/http"

	"github.com/labstack/gommon/log"
)

// errors package defines all the errors captured by the URL Shortner Application

var _ error = (*URLShortenerError)(nil)

type URLShortenerError struct {
	ErrorMessage string
	UserMessage  string
	HTTPCode     int
	Args         []string
}

func NewURLShortenerError(msg, userMsg string, code int, args ...string) *URLShortenerError {
	return &URLShortenerError{
		ErrorMessage: msg,
		UserMessage:  userMsg,
		HTTPCode:     code,
		Args:         args,
	}
}

func (u URLShortenerError) Error() string {
	return fmt.Sprintf("error_message: %s, %s, http_code: %d", u.ErrorMessage, u.Args, u.HTTPCode)
}

func SetError(errCode errorCode, args ...string) error {
	log.Info("Set Error ", "code ", errCode, "args ", args)

	switch errCode {
	case InternalServerError:
		return NewURLShortenerError("something crashed in the backend",
			"the server encountered an unexpected condition that prevented it from fulfilling the request.",
			http.StatusInternalServerError, args...)

	case InvalidURLError:
		return NewURLShortenerError("client error: invalid url inserted",
			"Invalid URL provided.",
			http.StatusBadRequest,
			args...,
		)

	case URLNotFoundError:
		return NewURLShortenerError("client error: shortened URL doesn't exist",
			"shortened URL is not registered with application.",
			http.StatusNotFound,
			args...,
		)

	case URLNotProvidedError:
		return NewURLShortenerError("required URL parameters missing",
			"Required parameter 'url' missing",
			http.StatusBadRequest,
			args...,
		)

	default:
		return NewURLShortenerError("undefinedError",
			"the server encountered an unexpected condition that prevented it from fulfilling the request.",
			http.StatusInternalServerError,
			args...,
		)

	}

}

type errorCode int

const (
	_ errorCode = iota
	InternalServerError
	InvalidURLError
	URLNotFoundError
	URLNotProvidedError
)
