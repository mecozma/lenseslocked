package main

import (
	"html/template"
	"os"
)

// Alternatively a named struct could be used.
type User struct {
	Name string
	Age  int
	Meta UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	// anonymous struct.
	// user := struct {
	// 	Name string
	// }{
	// 	Name: "Hank Hill",
	// }

	user := User{
		Name: "Hank Hill",
		Age:  112,
		Meta: UserMeta{
			Visits: 4,
		},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
