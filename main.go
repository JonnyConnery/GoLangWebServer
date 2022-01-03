package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func serve(w http.ResponseWriter, r *http.Request) {
	var index = "index.html"
	t, _ := template.ParseFiles(index)
	switch r.URL.Path {
	case "/":
		t.ExecuteTemplate(w, index, nil)
	case "/about":
		fmt.Fprint(w, "This is a web server built with Go")
	default:
		fmt.Fprint(w, "404")
	}
	fmt.Printf("Handling function with %s request\n", r.URL)
}
func main() {
	http.HandleFunc("/", serve)
	http.ListenAndServe("", nil)
}
