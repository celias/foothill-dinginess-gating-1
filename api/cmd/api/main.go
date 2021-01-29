package main

import (
	"simplesim/api/pkg/web"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	web.ConfigureRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
