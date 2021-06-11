package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/juliokscesar/unique-learningservice/unique-server/models"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "hello"}`))
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/api/", http.StatusFound)
	})

	router.HandleFunc("/api/", hello).Methods(http.MethodGet)

	// TESTS

	u, err := models.NewUser("Julio", "emailtest@gmail.com", "12345")
	if err != nil {
		log.Fatal("User error:", err)
	}

	c := models.NewCourse("CS50x", "Computer Science course from Harvard CS50")

	c.AppendStudents(u.ID)
	u.AppendCourses(c.ID)

	m := models.NewMaterial("Simple Material", "Material testing description", c.ID)

	a := models.NewAssignment("Problem Set Test", "New problem set testing", c.ID, time.Now().AddDate(0, 2, 0), false)

	router.HandleFunc("/api/test/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(u)
	}).Methods(http.MethodGet)

	router.HandleFunc("/api/test/course", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(c)
	}).Methods(http.MethodGet)

	router.HandleFunc("/api/test/material", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(m)
	}).Methods(http.MethodGet)

	router.HandleFunc("/api/test/assignment", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(a)
	}).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}
