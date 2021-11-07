package adapters

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"

	"github.com/DAGG3R09/url-shortener/service/errors"
)

type ServiceHTTPErrorResponse struct {
	UserMessage string `json:"message"`
}

func HttpErrorHandler(e error, c echo.Context) {
	fmt.Println("Eerrrrr", reflect.TypeOf(e))

	switch err := e.(type) {
	case *echo.HTTPError:
		c.JSON(err.Code, err.Message)
	case *errors.URLShortenerError:
		c.JSON(err.HTTPCode, MarshalToServiceHTTPResponse(err))
	default:
		c.JSON(http.StatusInternalServerError, "unexpected error"+e.Error())
	}
}

func MarshalToServiceHTTPResponse(e *errors.URLShortenerError) ServiceHTTPErrorResponse {
	return ServiceHTTPErrorResponse{UserMessage: e.UserMessage}
}
