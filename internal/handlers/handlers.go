package handlers

import (
	"html/template"
	"log"
	"net/http"

	"timetable/internal/mapper"
	"timetable/internal/togglclient"
)

var tableTemplate = template.Must(template.ParseFiles("templates/fragments/table.html"))

func HandleGetTable(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Toggl-Token")
	if token == "" {
		http.Error(w, "missing X-Toggl-Token header", http.StatusBadRequest)
		return
	}

	client := togglclient.New()

	projects, err := client.GetProjects(r.Context(), token)
	if err != nil {
		log.Printf("error fetching projects: %v", err)
		http.Error(w, "failed to fetch projects", http.StatusInternalServerError)
		return
	}
	log.Printf("projects: %+v", projects)

	entries, err := client.GetTimeEntries(r.Context(), token)
	if err != nil {
		log.Printf("error fetching time entries: %v", err)
		http.Error(w, "failed to fetch time entries", http.StatusInternalServerError)
		return
	}

	table := mapper.BuildTimeTable(entries, projects)

	w.Header().Set("Content-Type", "text/html")
	if err := tableTemplate.Execute(w, table); err != nil {
		log.Printf("error executing template: %v", err)
		http.Error(w, "failed to render template", http.StatusInternalServerError)
	}
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(table)

}
