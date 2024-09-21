package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

func RenderTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func ExecuteTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}
