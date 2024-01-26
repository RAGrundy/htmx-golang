package components

import (
	"github.com/labstack/echo/v4"
)

type WhoAmIComponent struct {
}

func WhoAmI(c echo.Context) error {
	return c.Render(200, "component/who-am-i.html", WhoAmIComponent{})
}
