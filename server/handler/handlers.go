package handler

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// errors.ParseErrors(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		// errors.ParseErrors(w, http.StatusMethodNotAllowed)
		return
	}
	// artists := grabjson.GetJsonData()
	html, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		// errors.ParseErrors(w, http.StatusInternalServerError)
		return
	}
	err = html.Execute(w, nil)
	if err != nil {
		// errors.ParseErrors(w, http.StatusInternalServerError)
		return
	}
}
