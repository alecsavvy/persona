package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	config, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	forwardClient := &http.Client{}

	app := NewApp(config, e, forwardClient)
	app.registerMiddleware()
	app.registerRoutes()
	err = app.serve()

	if err != nil {
		log.Fatal(err)
	}
}
