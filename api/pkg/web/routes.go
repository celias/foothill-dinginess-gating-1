package web

import "github.com/labstack/echo/v4"

// ConfigureRoutes configures the routes on the given echo instance.
func ConfigureRoutes(e *echo.Echo) {
	e.GET("/ping", handlePing)
	e.POST("/sim", handleSim)
}
