package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// LoadTemplates ...
func LoadTemplates(path string) {
	templates = template.Must(template.ParseGlob(path))
}

// ExecuteTemplate ...
func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}
