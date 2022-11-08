package handler

import (
	"groupie-tracker/server/grabjson"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Err(w, http.StatusNotFound)
		return
	}

	html, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}

	artists, err := grabjson.GetData()
	if err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}

	artists.FoundArtist = artists.Artist
	if err = html.Execute(w, &artists); err != nil {
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
		Err(w, http.StatusBadRequest)
		return
	}

	artist, err := grabjson.SearchArtist(text)
	if err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}

	if err = html.Execute(w, artist); err != nil { // artist it's search result
		Err(w, http.StatusInternalServerError)
		return
	}
}

func isTextEmpty(s string) bool {
	if s == "" {
		return true
	}
	return false
}
