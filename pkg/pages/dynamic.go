package pages

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

type DynamicPage struct {
	Name string
}

func DynamicPageRouter(c echo.Context) error {
	p := removePages(c.Request().URL.Path)

	switch p {
	case "/":
		return c.Render(200, "index.html", DynamicPage{Name: "Index"})
	default:
		err := c.Render(200, fmt.Sprintf("%s.html", strings.TrimLeft(p, "/")), DynamicPage{Name: strings.TrimLeft(p, "/")})
		if err != nil {
			return c.Render(404, "notfound.html", DynamicPage{})
		}
		return err
	}
}

func removePages(rp string) (p string) {
	p = rp

	p = strings.TrimPrefix(p, "/global")
	p = strings.TrimPrefix(p, "/index")
	p = strings.TrimPrefix(p, "/notfound")
	p = strings.TrimPrefix(p, "/servererror")
	return
}
