package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func handleRequests(w http.ResponseWriter, r *http.Request) {

	var filePath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Jira Backlog Converter",
		"name":  "Testing",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handleRequests)
	http.Handle("/assets/",
		http.StripPrefix("/assets/",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("Server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
