package main

import (
	"html/template"
	"os"
)

// Alternatively a named struct could be used.
type User struct {
	Name string
	Bio  string
	Age  int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Hank Hill",
		Bio:  `<script>alert("HAHA, careful there!");</script>`,
		Age:  112,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}