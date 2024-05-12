package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

//go:embed index.html block/film.tmpl
var files embed.FS

func getTemplate() (*template.Template, error) {
	return template.ParseFS(files, "index.html", "block/film.tmpl")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(getTemplate())
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}
		t.Execute(w, films)
	})

	mux.HandleFunc("POST /add-film/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		t := template.Must(getTemplate())
		t.ExecuteTemplate(w, "film", Film{Title: title, Director: director})
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
