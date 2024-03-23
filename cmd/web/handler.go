package main

import (
	"html/template"
	"log"
	"net/http"
)

func viewApp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	}

	ts, err := getHtmlTemplate(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	query := r.URL.Query().Get("q")
	if query == "" {
		createPage(w, ts)
	} else {
		createQueryPage(w, ts, query)
	}
}

func getHtmlTemplate(w http.ResponseWriter) (*template.Template, error) {
	filesPathes := []string{
		"./ui/html/base.html",
		"./ui/html/partials/footer.html",
		"./ui/html/pages/home.html",
	}
	ts, err := template.ParseFiles(filesPathes...)
	return ts, err
}

func createPage(w http.ResponseWriter, ts *template.Template) {
	err := ts.ExecuteTemplate(w, "base", nil)
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

	err = ts.ExecuteTemplate(w, "base", cssPropData)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Error Appeared", http.StatusInternalServerError)
	}
}
