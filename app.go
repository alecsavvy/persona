package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type App struct {
	config        *Config
	server        *echo.Echo
	forwardClient *http.Client
}

func NewApp(config *Config, server *echo.Echo, forwardClient *http.Client) *App {
	return &App{
		config, server, forwardClient,
	}
}

func (app *App) forwardRequest(c echo.Context) error {
	originalRequest := c.Request()

	req, err := http.NewRequest(originalRequest.Method, app.config.IdentityUrl, originalRequest.Body)
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

func (app *App) registerRoutes() {
	app.server.POST("/authentication", app.forwardRequest)
	app.server.GET("/authentication", app.forwardRequest)
	app.server.GET("/authentication/check", app.forwardRequest)
	app.server.GET("/cognito_signature", app.forwardRequest)
	app.server.POST("/cognito_webhook/flow", app.forwardRequest)
	app.server.POST("/cognito_retry/:handle", app.forwardRequest)
	app.server.GET("/cognito_recent_exists/:handle", app.forwardRequest)
	app.server.POST("/cognito_flow", app.forwardRequest)
	app.server.POST("/eth_relay", app.forwardRequest)
	app.server.GET("/eth_relayer", app.forwardRequest)
	app.server.POST("/fp/webhook", app.forwardRequest)
	app.server.GET("/fp", app.forwardRequest)
	app.server.GET("/fp/counts/:userId", app.forwardRequest)
	app.server.GET("/health_check/relay", app.forwardRequest)
	app.server.GET("/health_check", app.forwardRequest)
	app.server.GET("/health_check/poa", app.forwardRequest)
	app.server.GET("/balance_check", app.forwardRequest)
	app.server.GET("/eth_balance_check", app.forwardRequest)
	app.server.GET("/sol_balance_check", app.forwardRequest)
	app.server.GET("/notification_check", app.forwardRequest)
	app.server.GET("/db_check", app.forwardRequest)
	app.server.GET("/rewards_check", app.forwardRequest)
	app.server.GET("/id_signals", app.forwardRequest)
	app.server.POST("/record_ip", app.forwardRequest)
	app.server.POST("/instagram", app.forwardRequest)
	app.server.POST("/instagram/associate", app.forwardRequest)
	app.server.GET("/location", app.forwardRequest)
	app.server.GET("/notifications", app.forwardRequest)
	app.server.POST("/notifications", app.forwardRequest)
	app.server.POST("/notifications/all", app.forwardRequest)
	app.server.POST("/notifications/clear_badges", app.forwardRequest)
	app.server.POST("/notifications/settings", app.forwardRequest)
	app.server.GET("/notifications/settings", app.forwardRequest)
	app.server.POST("/announcements", app.forwardRequest)
	app.server.POST("/notifications/subscription", app.forwardRequest)
	app.server.GET("/notifications/subscription", app.forwardRequest)
	app.server.GET("/push_notifications/settings", app.forwardRequest)
	app.server.POST("/push_notifications/settings", app.forwardRequest)
	app.server.POST("/push_notifications/device_token", app.forwardRequest)
	app.server.POST("/push_notifications/device_token/deregister", app.forwardRequest)
	app.server.GET("/push_notifications/device_token/enabled", app.forwardRequest)
	app.server.GET("/push_notifications/browser/settings", app.forwardRequest)
	app.server.POST("/push_notifications/browser/settings", app.forwardRequest)
	app.server.GET("/push_notifications/browser/enabled", app.forwardRequest)
	app.server.POST("/push_notifications/browser/register", app.forwardRequest)
	app.server.POST("/push_notifications/browser/deregister", app.forwardRequest)
	app.server.POST("/push_notifications/safari/:version/pushPackages/:websitePushID", app.forwardRequest)
	app.server.POST("/push_notifications/safari/:version/devices/:deviceToken/registrations/:websitePushID", app.forwardRequest)
	app.server.DELETE("/push_notifications/safari/:version/devices/:deviceToken/registrations/:websitePushID", app.forwardRequest)
	app.server.POST("/push_notifications/safari/:version/log", app.forwardRequest)
	app.server.POST("/reactions", app.forwardRequest)
	app.server.GET("/reactions", app.forwardRequest)
	app.server.POST("/recovery", app.forwardRequest)
	app.server.POST("/relay", app.forwardRequest)
	app.server.GET("/scores", app.forwardRequest)
	app.server.POST("/score/hcaptcha", app.forwardRequest)
	app.server.GET("/social_handles", app.forwardRequest)
	app.server.POST("/social_handles", app.forwardRequest)
	app.server.POST("/stripe/session", app.forwardRequest)
	app.server.GET("/tiktok", app.forwardRequest)
	app.server.OPTIONS("/tiktok/access_token", app.forwardRequest)
	app.server.POST("/tiktok/access_token", app.forwardRequest)
	app.server.POST("/tiktok/associate", app.forwardRequest)
	app.server.GET("/tracks/listen/solana/status", app.forwardRequest)
	app.server.POST("/tracks/:id/listen", app.forwardRequest)
	app.server.GET("/tracks/history", app.forwardRequest)
	app.server.POST("/tracks/listens/:timeframe*?", app.forwardRequest)
	app.server.GET("/tracks/listens/:timeframe*?", app.forwardRequest)
	app.server.POST("/tracks/trending/:time*?", app.forwardRequest)
	app.server.GET("/tracks/trending/:time*?", app.forwardRequest)
	app.server.GET("/users/listens/top", app.forwardRequest)
	app.server.GET("/users/listens", app.forwardRequest)
	app.server.GET("/listens/bulk", app.forwardRequest)
	app.server.POST("/twitter", app.forwardRequest)
	app.server.POST("/twitter/callback", app.forwardRequest)
	app.server.GET("/twitter/handle_lookup", app.forwardRequest)
	app.server.POST("/twitter/associate", app.forwardRequest)
	app.server.POST("/user", app.forwardRequest)
	app.server.GET("/users/check", app.forwardRequest)
	app.server.POST("/users/update", app.forwardRequest)
	app.server.PUT("/user/email", app.forwardRequest)
	app.server.GET("/user/email", app.forwardRequest)
	app.server.POST("/user/associate", app.forwardRequest)
	app.server.GET("/auth_migration", app.forwardRequest)
	app.server.POST("/auth_migration", app.forwardRequest)
	app.server.POST("/transaction_metadata", app.forwardRequest)
	app.server.GET("/transaction_metadata", app.forwardRequest)
	app.server.GET("/userEvents", app.forwardRequest)
	app.server.POST("/userEvents", app.forwardRequest)
	app.server.POST("/user_playlist_favorites", app.forwardRequest)
	app.server.GET("/user_playlist_favorites", app.forwardRequest)
	app.server.GET("/user_playlist_updates", app.forwardRequest)
	app.server.POST("/user_playlist_updates", app.forwardRequest)
	app.server.POST("/email/welcome", app.forwardRequest)
	app.server.POST("/wormhole_relay", app.forwardRequest)
}

func (app *App) serve() error {
	return app.server.Start(":7000")
}
