package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	logger := InitLogger()

	config, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	logger.Info("configuration", "env", config.Environment, "identity", config.IdentityUrl)

	e := echo.New()

	forwardClient := &http.Client{}

	app := NewApp(logger, config, e, forwardClient)
	err = app.run()

	if err != nil {
		log.Fatal(err)
	}
}
