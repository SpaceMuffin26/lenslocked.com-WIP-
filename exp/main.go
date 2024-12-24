package main

import (
	"html/template"
	"net/http"
)

var homeTemplate *template.Template

func main() {
	homeTemplate, _ = template.ParseFiles("views/home.gohtml")
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}
