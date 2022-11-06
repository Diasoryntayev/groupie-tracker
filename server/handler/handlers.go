package handler

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Err(w, http.StatusNotFound)
		return
	}
	// artists := grabjson.GetJsonData()
	html, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}
	if err = html.Execute(w, nil); err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}
}

func Artist(w http.ResponseWriter, r *http.Request) {
	// detailArtist := grabjson.GetArtistById(id)

	html, err := template.ParseFiles("./ui/html/details.html")
	if err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}
	if err = html.Execute(w, nil); err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}
	text := r.FormValue("search")
	if isTextEmpty(text) {
		return
	}

	if err = html.Execute(w, artist); err != nil { // artist it's search result
		Err(w, http.StatusInternalServerError)
		return
	}
}
