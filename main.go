package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mecozma/lenslocked/controllers"
	"github.com/mecozma/lenslocked/templates"
	"github.com/mecozma/lenslocked/views"
)

func main() {
	r := chi.NewRouter()

	tpl := views.Must(views.ParseFS(
		templates.FS,
		"home.gohtml",
		"tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(
		templates.FS,
		"contact.gohtml",
		"tailwind.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(
		templates.FS,
		"faq.gohtml",
		"tailwind.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))

	tpl = views.Must(views.ParseFS(
		templates.FS,
		"gallery.gohtml",
		"tailwind.gohtml"))
	r.Get("/gallery", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml",
		"tailwind.gohtml"))
	r.Get("/signup", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Our server is a bit shy now, please try again :)",
			http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
