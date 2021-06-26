package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/juliokscesar/unique-learningservice/unique-server/controller"
	"github.com/juliokscesar/unique-learningservice/unique-server/models"
	"github.com/juliokscesar/unique-learningservice/unique-server/utils"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "hello"}`))
}

func main() {
	router := mux.NewRouter()

	err := controller.ControllerInit()
	if err != nil {
		log.Fatal("Controller initalizing errror:", err)
	}

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/api/", http.StatusFound)
	})

	router.HandleFunc("/api/", hello).Methods(http.MethodGet)

	/* TEST ROUTERS */

	u, err := models.NewUser("Julio", "emailtest@gmail.com", "12345")
	if err != nil {
		log.Fatal("User error:", err)
	}
	err = controller.RegisterUser(u)
	if err != nil {
		log.Println("Inserting user error:", err)
	}

	c := models.NewCourse("CS50x", "Computer Science course from Harvard CS50")

	c.AppendStudents(u.ID)
	u.AppendCourses(c.ID)

	err = controller.InsertOneCourse(c)
	if err != nil {
		log.Println(err)
	}

	m := models.NewMaterial("Simple Material", "Material testing description", c.ID)

	err = controller.InsertOneMaterial(m)
	if err != nil {
		log.Println(err)
	}

	a := models.NewAssignment("Problem Set Test", "New problem set testing", c.ID, time.Now().AddDate(0, 2, 0), false)

	err = controller.InsertOneAssignment(a)
	if err != nil {
		log.Println(err)
	}

	router.HandleFunc("/api/test/user", func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(u)
	}).Methods(http.MethodGet)

	router.HandleFunc("/api/test/user/id", func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)

		foundUser, err := controller.GetUserFromID(u.ID.Hex())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(foundUser)
	})

	router.HandleFunc("/api/test/course", func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(c)
	}).Methods(http.MethodGet)

	router.HandleFunc("/api/test/course/id", func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)

		foundCourse, err := controller.GetCourseFromId(c.ID.Hex())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(foundCourse)
	})

	router.HandleFunc("/api/test/material", func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(m)
	}).Methods(http.MethodGet)

	router.HandleFunc("/api/test/material/id", func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)

		foundMaterial, err := controller.GetMaterialFromId(m.ID.Hex())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(foundMaterial)
	})

	router.HandleFunc("/api/test/assignment", func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(a)
	}).Methods(http.MethodGet)

	router.HandleFunc("/api/test/assignment/id", func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)

		foundAssignment, err := controller.GetAssignmentFromId(a.ID.Hex())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(foundAssignment)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
