package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/index.html", "templates/header.html")
	http.FileServer(http.Dir("static/"))
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	temp.ExecuteTemplate(w, "index", nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/create.html", "templates/header.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	temp.ExecuteTemplate(w, "create", nil)
}

func handleFunc() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	// http.HandleFunc("/create_release", create_release)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}
