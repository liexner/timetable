package main

import (
	"log"
	"net/http"
	"text/template"
	"timetable/internal/handlers"
)

func main() {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("GET /table", handlers.HandleGetTable)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		tmpl.Execute(w, nil)
	})

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
