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
	fmt.Printf("\n----------------------------------------\n")
	fmt.Printf(tmpl)
	fmt.Printf("\n----------------------------------------\n")
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

//Load the page
func loadPage(title string) (*Page, error) {
	return &Page{Title: title}, nil
}

//Http handler
func handler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[1:]
	p, err := loadPage(title)

	if err != nil {
		http.Redirect(w, r, "404", http.StatusFound)
		return
	}

	renderTemplate(w, "view/" + title, p)
}

//Main
func main() {
	http.HandleFunc("/", handler)

	fmt.Printf("Listening in the port :4000")
	http.ListenAndServe(":4000", nil)
}
