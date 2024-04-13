package main

import (
	"html/template"
	"net/http"
)

func createTemplate(w http.ResponseWriter, apply func(w http.ResponseWriter, ts *template.Template)) {
	filesPathes := []string{
		"./ui/html/base.html",
		"./ui/html/partials/footer.html",
		"./ui/html/pages/home.html",
		"./ui/html/forms/search.html",
	}
	ts, err := template.ParseFiles(filesPathes...)
	if err != nil {
		http.Error(w, "HTML Creating error", http.StatusInternalServerError)
		return
	}
	apply(w, ts)
}
