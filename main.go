package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site that now refreshes dynamically!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:me@gmail.com\">
	me@gmail.com</a>.`)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<h1>FAQ page</h1><h2>Is there a free version?</h2><p>Yes! We offer a free trial for 
	30 days on any paid plans.</p><h2>What are your support hours?</h2><p> We have support staff 
	answering emails 24/7, though response times may be a bit slower on weekends.</p><h2>How do I contact
	support?</h2><p>Email us - <a href="support@lenslocked.com">support@lenslockes.com</a></p>`)
}

func galleryHandler(w http.ResponseWriter, r *http.Request) {
	userName := chi.URLParam(r, "userName")
	w.Write([]byte(fmt.Sprintf("Hello %s!", userName)))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/gallery/{userName}", galleryHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Our server is a bit shy now, please try again :)", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
