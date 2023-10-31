package main

import (
	"html/template"
	"net/http"
)

var pageTemplate *template.Template

func main() {
	mux := http.NewServeMux()
	pageTemplate = template.Must(template.ParseFiles("index.html"))
	mux.Handle("/", http.FileServer(http.Dir("/")))
	http.ListenAndServe(":8080", mux)
}
