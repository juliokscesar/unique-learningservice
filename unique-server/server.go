package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "hello"}`))
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}
