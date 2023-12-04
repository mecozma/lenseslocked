package main

import (
	"fmt"
	"net/http"
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

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "The server is a bit shy today, try again :)", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router)
}
