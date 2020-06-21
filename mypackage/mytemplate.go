package mypackage

import (
	"io"
  	"html/template"
  	"github.com/labstack/echo/v4"
)

type Template struct {
    templates *template.Template
}

func (template *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return template.templates.ExecuteTemplate(w, name, data)
}

func GetRenderer() *Template {
	return &Template{
	    templates: template.Must(template.ParseGlob("views/*.html")),
	}
}