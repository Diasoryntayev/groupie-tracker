package handler

import (
	"net/http"
	"text/template"
)

type ErrorHTTP struct {
	Status  int
	Message string
}

func Err(w http.ResponseWriter, status int) {
	Err := ErrorHTTP{status, http.StatusText(status)}
	temp, err := template.ParseFiles("./ui/html/error.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(Err.Status)
	if err = temp.Execute(w, Err); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
