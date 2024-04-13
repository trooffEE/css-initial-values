package main

import (
	"html/template"
	"log"
	"net/http"
)

func richHtmlWithData(w http.ResponseWriter, ts *template.Template, data any) {
	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Error Appeared", http.StatusInternalServerError)
	}
}

func createPage(w http.ResponseWriter, ts *template.Template, query string) {
	cssPropData, err := ScrappedCSSFormalDefinition{}, error(nil)
	if query != "" {
		cssPropData, err = getCssInitialDataByQuery(query)
	}
	if err != nil {
		http.Error(w, "Scrapping implementation error", http.StatusInternalServerError)
		return
	}

	richHtmlWithData(w, ts, cssPropData)
}

func handleRedirect(w http.ResponseWriter, r *http.Request, query string) {
	url := "/"
	if query != "" {
		url += "?query=" + query
	}

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func viewApp(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")
	if r.URL.Path != "/" {
		handleRedirect(w, r, query)
		return
	}

	createTemplate(w, func(w http.ResponseWriter, ts *template.Template) {
		createPage(w, ts, query)
	})
}
