package web

import (
	"github.com/labstack/echo/v4"
)

func handlePing(c echo.Context) error {
	return c.JSON(200, "pong")
}

func handleSim(c echo.Context) error {
	// FIXME: Call sim.Sim() with the POST'd input
	return c.JSON(
		501,
		map[string]string{
			"msg": "not implemented",
		},
	)
}
