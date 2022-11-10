package server

import (
	"groupie-tracker/server/handler"
	"log"
	"net/http"
)

func Server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", GetZapros(handler.Home))
	mux.HandleFunc("/details", GetZapros(handler.Artist))
	mux.HandleFunc("/searched/", GetZapros(handler.SearchHandler))

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Запуск веб-сервера на http://127.0.0.1:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func GetZapros(handleFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			handler.Err(w, http.StatusMethodNotAllowed)
			return
		}
		handleFunc(w, r)
	}
}
