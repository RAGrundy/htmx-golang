package main

import (
	"html/template"
	"io"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/ragrundy/htmx-go/pkg/pages"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	tmpls, err := template.New("").ParseGlob("web/templates/*/*.html")
	if err != nil {
		log.Fatalf("couldn't initialize templates: %v", err)
	}

	e.Renderer = &TemplateRenderer{
		templates: tmpls,
	}

	e.Use(middleware.Logger())

	e.Static("/scripts", "web/js")
	e.Static("/css", "web/css")

	e.GET("/", pages.Index)

	e.Logger.Fatal(e.Start(":1323"))

}