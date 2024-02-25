package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary forwards requests to identity service
// @Description handles any route and method and forwards it to the identity service as if the caller called identity directly
// @Accept json
// @Produce json
// @Success 200
// @Router / [get]
// @Router / [post]
func (app *App) forwardRequest(c echo.Context) error {
	originalRequest := c.Request()
	uri := originalRequest.RequestURI
	forwardedUrl := fmt.Sprintf("%s%s", app.config.IdentityUrl, uri)

	app.logger.Info("forwarding request", "url", forwardedUrl, "method", originalRequest.Method)

	req, err := http.NewRequest(originalRequest.Method, forwardedUrl, originalRequest.Body)
	if err != nil {
		return err
	}

	for name, values := range originalRequest.Header {
		for _, value := range values {
			req.Header.Add(name, value)
		}
	}

	// Send the request
	resp, err := app.forwardClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	c.Response().Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	return c.Stream(resp.StatusCode, resp.Header.Get("Content-Type"), resp.Body)
}
