package handler

import (
	"groupie-tracker/server/grabjson"
	"html/template"
	"net/http"
	"strconv"
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
	html, err := template.ParseFiles("./ui/html/details.html")
	if err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}

	artists, err := grabjson.GetData()
	if err != nil {
		Err(w, http.StatusInternalServerError)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		Err(w, http.StatusNotFound)
		return
	}
	if id < 1 || id > len(artists.Artist) {
		Err(w, http.StatusNotFound)
		return
	}

	if err = html.Execute(w, &artists.Artist[id-1]); err != nil {
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
