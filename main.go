package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "fatal error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	logger := InitLogger()

	config, err := NewConfig()
	if err != nil {
		return err
	}

	logger.Info("configuration", "env", config.Environment, "identity", config.IdentityUrl)

	e := echo.New()

	forwardClient := &http.Client{}

	app := NewApp(logger, config, e, forwardClient)
	err = app.run()

	if err != nil {
		return err
	}

	return nil
}
