package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseFiles("templates/index.html"))

func renderTemplate(w http.ResponseWriter, filepath string) {
	err := templates.ExecuteTemplate(w, filepath, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	if path == "/" {
		http.Redirect(w, req, "/index.html", http.StatusFound)
	} else if path != "/index.html" {
		http.NotFound(w, req)
	} else {
		log.Println(path)
		// _, cwd := os.Getwd()
		// log.Println(cwd)
		renderTemplate(w, "index.html")
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
