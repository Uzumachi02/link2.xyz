package utils

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// Template ...
type Template struct {
	templates *template.Template
}

// Render ...
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// NewTemplate ...
func NewTemplate(glob string) *Template {
	return &Template{
		templates: template.Must(template.ParseGlob(glob)),
	}
}
