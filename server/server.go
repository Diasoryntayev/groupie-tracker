package server

import (
	"log"
	"net/http"
)

func Server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/details", artist)

	log.Println("Запуск веб-сервера на http://127.0.0.1:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
