package web

import (
	"simplesim/api/pkg/sim"

	"github.com/labstack/echo/v4"
)

func handlePing(c echo.Context) error {
	return c.JSON(200, "pong")
}

func handleSim(c echo.Context) error {
	// FIXME: Call sim.Sim() with the POST'd input
	resp := sim.Results{}
	return c.JSON(200, resp)
}
