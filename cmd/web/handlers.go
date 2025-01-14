package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./ui/html/pages/home.tmpl.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, "home")
	if err != nil {
		log.Fatal(err)
	}
}
