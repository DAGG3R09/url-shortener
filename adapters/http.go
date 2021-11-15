package adapters

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/DAGG3R09/url-shortener/service"
)

func NewHTTPRouter(s *service.URLShortnerService) *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.HTTPErrorHandler = HttpErrorHandler
	e.GET("/", running)

	addUrlShortenerRoutes(e, s)
	return e
}

func addUrlShortenerRoutes(e *echo.Echo, s *service.URLShortnerService) {
	eg := e.Group(shortnerPrefix)

	eg.POST("/", shortenURL(s))
	eg.POST("", shortenURL(s))
	eg.GET("/:url", redirectURLHandler(s))
	eg.GET("/:url/", redirectURLHandler(s))
}

func running(ctx echo.Context) error {
	return ctx.JSON(200, "we are up and running")
}

func shortenURL(s *service.URLShortnerService) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		params := struct {
			URL string `json:"url"`
		}{}

		err := ctx.Bind(&params)
		if err != nil {
			return err
		}

		resp := struct {
			OriginalURL string `json:"original_url"`
			ShortURL    string `json:"short_url"`
		}{
			OriginalURL: params.URL,
		}

		resp.ShortURL, err = s.Shorten(params.URL)
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusCreated, resp)
	}
}

func redirectURLHandler(s *service.URLShortnerService) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		shortURL := ctx.Param("url")

		redirectURL, err := s.FetchRedirectURL(shortURL)
		if err != nil {
			return err
		}

		return ctx.Redirect(http.StatusTemporaryRedirect, redirectURL)
	}
}

const (
	shortnerPrefix = simplePrefix

	noPrefix     = ""
	simplePrefix = "/url-shortener"
)
