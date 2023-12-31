package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("hello world")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "title_a", Director: "director_a"},
				{Title: "title_b", Director: "director_b"},
				{Title: "title_c", Director: "director_c"},
			},
		}
		template.Execute(w, films)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
