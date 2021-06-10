package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	//"github.com/juliokscesar/unique-learningservice/unique-server/uniquedb/models"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "hello"}`))
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "hello from post"}`))
	}).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", router))
}
