package setup

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/ragrundy/htmx-golang/pkg/pages"
	"github.com/ragrundy/htmx-golang/pkg/susurrus"
)

func SetupTemplates(e *echo.Echo) {
	ltmpls := susurrus.AddFilesInDirectoryToTemplateWithPrefix(e,
		susurrus.TemplateDirectory{Directory: "web/templates/components", Prefix: "components"},
		susurrus.TemplateDirectory{Directory: "web/templates/global", Prefix: "global"},
		susurrus.TemplateDirectory{Directory: "web/templates/pages"})

	fmt.Println(ltmpls)
}

func SetupStatic(e *echo.Echo) {
	e.Static("/scripts", "web/js")
	susurrus.AddStaticFilesInDirectory(e,
		susurrus.StaticDirectory{Directory: "web/templates", Suffix: ".js", Route: "/scripts"},
		susurrus.StaticDirectory{Directory: "web/templates", Suffix: ".css", Route: "/css"})
}

func SetupGetRoutes(e *echo.Echo) {
	e.GET("/*", pages.DynamicPageRouter)
}

func SetupPostRoutes(e *echo.Echo) {
}
