package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Alternatively a named struct could be used.
type User struct {
	Name      string
	Bio       Biography
	Age       int
	Interests Interest
}

type Biography struct {
	Description string
}

type Interest struct {
	Hobby []string
}

// homeHandler function handles the home page.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// tplPath := filepath.Join("cmd", "exp", "templates", "home.gohtml")
	user := User{
		// Name: "Hank Hill",
		// Age: 112,
		Bio: Biography{
			Description: `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been
			 the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley
			 of type and scrambled it to make a type specimen book. It has survived not only five centuries,
			 but also the leap into electronic typesetting, remaining essentially unchanged`,
		},
		// Interests: Interest{
		// 	Hobby: []string{"Golang", "webapps", "wintage cars"},
		// },
	}

	tpl, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		log.Printf("Parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, user)
	if err != nil {
		log.Printf("Executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Our server is a bit shy now, please try again :)", http.StatusNotFound)
	})

	fmt.Println("Starting the experimental server on :3000...")
	http.ListenAndServe(":3000", r)
}
