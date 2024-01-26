package pages

import (
	"fmt"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

type DynamicPage struct {
}

func Dynamic(c echo.Context) error {
	p := c.Request().URL.Path
	log.Printf("dynamic routing: %s", p)
	switch p {
	case "/":
		return c.Render(200, "index.html", DynamicPage{})
	default:
		return c.Render(200, fmt.Sprintf("%s.html", strings.TrimLeft(p, "/")), DynamicPage{})
	}
}
