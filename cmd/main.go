package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/ragrundy/htmx-golang/internal/app/setup"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	setup.SetupTemplates(e)
	setup.SetupStatic(e)
	setup.SetupGetRoutes(e)
	setup.SetupPostRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
