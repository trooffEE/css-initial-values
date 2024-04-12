package main

import (
	"html/template"
	"log"
	"net/http"
)

func createPage(w http.ResponseWriter, ts *template.Template, data any) {
	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Error Appeared", http.StatusInternalServerError)
	}
}

func createQueryPage(w http.ResponseWriter, ts *template.Template, query string) {
	cssPropData, err := getCssInitialDataByQuery(query)
	if err != nil {
		http.Error(w, "Scrapping implementation error", http.StatusInternalServerError)
		return
	}

	createPage(w, ts, cssPropData)
}

func viewApp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	}

	createHTML(w, func(w http.ResponseWriter, ts *template.Template) {
		query := r.URL.Query().Get("q")
		if query == "" {
			createPage(w, ts, nil)
		} else {
			createQueryPage(w, ts, query)
		}
	})
}
