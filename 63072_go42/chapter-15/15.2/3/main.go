package main

import (
	"net/http"
	"text/template"
)

func tHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("*.html.tmpl"))
	t.Execute(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", tHandler)
	http.ListenAndServe(":8080", nil)
}
