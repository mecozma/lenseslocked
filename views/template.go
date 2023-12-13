package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

// ParseFS function parses the html templates.
func ParseFS(fs fs.FS, pattern string) (Template, error) {
	tpl, err := template.ParseFS(fs, pattern)
	if err != nil {
		return Template{}, fmt.Errorf("parseFS template: %w", err)
	}
	t := Template{
		htmlTpl: tpl,
	}
	return t, nil
}

type Template struct {
	htmlTpl *template.Template
}

// Execute function renders the html templates.
func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, nil) //TODO pass `data` to Execute()
	if err != nil {
		log.Printf("Executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}
