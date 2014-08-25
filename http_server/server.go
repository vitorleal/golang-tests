package main

import (
	"html/template"
	"fmt"
	"net/http"
)

//Page struct
type Page struct {
	Title string
}

//Render the template
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

//Http handler
func handler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[1:]
	page := &Page{Title: title}

	renderTemplate(w, "view/index", page)
}

//Not found page
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	page := &Page{Title: "Not found"}
	renderTemplate(w, "view/404", page)
}

//Main
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/404", notFoundHandler)

	fmt.Printf("Listening in the port :4000")
	http.ListenAndServe(":4000", nil)
}
