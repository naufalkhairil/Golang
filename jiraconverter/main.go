package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/naufalkhairil/jiraconverter/modules"
)

func routeIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var tmpl = template.Must(template.New("form").ParseFiles("views/view.html"))
		var err = tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "ErrorGet", http.StatusBadRequest)
}

func routeSubmitPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tmpl = template.Must(template.New("results").ParseFiles("views/view.html"))

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var backlog = r.FormValue("input_backlog")

		full_string := []string{}
		split_backlog := strings.Split(backlog, "\n")

		for i := 0; i < len(split_backlog)/2; i++ {
			das_res, err := modules.FindDAS(split_backlog[(i*2)+1])
			if err != nil {
				fmt.Println(err)
			}

			das_rep := modules.ReplaceString(split_backlog[i*2])

			das_assign, err := modules.FindAssignee(split_backlog[(i*2)+1])
			if err != nil {
				fmt.Println(err)
			}

			concat_str := strings.Join([]string{das_res, das_rep, das_assign}, " ")
			full_string = append(full_string, concat_str)
		}
		// var project_name = r.FormValue("input_project_name")

		var data = map[string][]string{"backlog": full_string}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "ErrorPost", http.StatusBadRequest)

}

func main() {
	http.HandleFunc("/", routeIndexGet)
	http.HandleFunc("/process", routeSubmitPost)
	http.Handle("/assets/",
		http.StripPrefix("/assets/",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("Server started at http://localhost:9000")
	http.ListenAndServe(":9000", nil)
}
