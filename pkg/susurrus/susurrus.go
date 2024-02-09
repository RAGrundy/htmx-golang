package susurrus

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type TemplateDirectory struct {
	Directory string
	Prefix    string
}

func AddFilesInDirectoryToTemplateWithPrefix(e *echo.Echo, pts ...TemplateDirectory) {
	tmpls := template.New("")

	for _, pt := range pts {
		err := filepath.Walk(pt.Directory,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if strings.HasSuffix(path, ".html") {
					rf, err := os.ReadFile(path)
					if err != nil {
						log.Println(err)
						return err
					}
					rfs := string(rf)

					var sb strings.Builder
					if pt.Prefix != "" {
						sb.WriteString(fmt.Sprintf("%s/", pt.Prefix))
					}

					uPath := strings.ReplaceAll(path, "\\", "/")
					uPath = strings.ReplaceAll(uPath, fmt.Sprintf("%s/", strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))), "")
					sb.WriteString(strings.TrimPrefix(uPath, pt.Directory+"/"))

					fmt.Println(sb.String())
					tmpls.New(sb.String()).Parse(rfs)
				}
				return nil
			})
		if err != nil {
			log.Println(err)
		}
	}

	e.Renderer = &TemplateRenderer{
		templates: tmpls,
	}
}

type StaticDirectory struct {
	Directory string
	Suffix    string
	Route     string
}

func AddStaticFilesInDirectory(e *echo.Echo, sds ...StaticDirectory) {
	for _, sd := range sds {
		err := filepath.Walk(sd.Directory,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if strings.HasSuffix(path, sd.Suffix) {
					e.File(fmt.Sprintf("%s/%s", sd.Route, info.Name()), path)
				}
				return nil
			})
		if err != nil {
			log.Println(err)
		}
	}
}
