package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Parse function parses the html templates.
func Parse(filepath string) (Template, error) {
	// tplPath joines the path before passing it to Parsefiles method to ensure it is working on any OS.
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
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
