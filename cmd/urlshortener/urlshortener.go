package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/DAGG3R09/url-shortener/adapters"
	"github.com/DAGG3R09/url-shortener/repository"
	"github.com/DAGG3R09/url-shortener/service"
	"github.com/DAGG3R09/url-shortener/service/shorteners"
)

func main() {

	u := service.NewURLShortnerService(
		shorteners.NewBase62Shortener(),
		repository.NewMapRepository(),
		service.MinValueBase62Number, service.MaxValueBase62Number,
	)

	e := adapters.NewHTTPRouter(u)

	e.Use(middleware.Logger())

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT)
	go func() {
		<-shutdown

		log.Info("Shutting down web server...")

		err := e.Shutdown(context.Background())
		if err != nil {
			log.Error("error while shutting down server", err)
		}

		os.Exit(1)
	}()

	err := e.Start(":8080")
	if err != nil {
		log.Error("Failed to start server", err)
		return
	}

}
