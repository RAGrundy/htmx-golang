package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/ragrundy/htmx-golang/pkg/components"
	"github.com/ragrundy/htmx-golang/pkg/pages"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	log.Printf("Render template: %s", name)
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	tmpls := template.New("")

	f, err := os.Open("web/templates/component")
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range files {
		fmt.Println(v.Name(), v.IsDir())
		f, err := os.ReadFile(fmt.Sprintf("web/templates/component/%s", v.Name()))
		if err != nil {
			fmt.Print(err)
		}
		str := string(f)

		tmpls.New(fmt.Sprintf("component/%s", v.Name())).Parse(str)
	}

	f, err = os.Open("web/templates/global")
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err = f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range files {
		fmt.Println(v.Name(), v.IsDir())
		f, err := os.ReadFile(fmt.Sprintf("web/templates/global/%s", v.Name()))
		if err != nil {
			fmt.Print(err)
		}
		str := string(f)

		tmpls.New(fmt.Sprintf("global/%s", v.Name())).Parse(str)
	}

	f, err = os.Open("web/templates/page")
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err = f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range files {
		fmt.Println(v.Name(), v.IsDir())
		f, err := os.ReadFile(fmt.Sprintf("web/templates/page/%s", v.Name()))
		if err != nil {
			fmt.Print(err)
		}
		str := string(f)

		tmpls.New(fmt.Sprintf("%s", v.Name())).Parse(str)
	}

	log.Printf("%s", tmpls.DefinedTemplates())

	e.Renderer = &TemplateRenderer{
		templates: tmpls,
	}

	e.Static("/scripts", "web/js")
	e.Static("/css", "web/css")

	e.GET("/*", pages.Dynamic)
	e.GET("/components/who-am-i", components.WhoAmI)

	e.Logger.Fatal(e.Start(":1323"))
}
