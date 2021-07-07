package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/juliokscesar/unique-learningservice/unique-server/controller"
)

func main() {
	router := mux.NewRouter()

	err := controller.ControllerInit()
	if err != nil {
		log.Fatal("Controller initalizing errror:", err)
	}

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/api/", http.StatusFound)
	})

	router.HandleFunc("/api/user/register", controller.RegisterUserHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/api/user/login", controller.LoginUserHandler).Methods(http.MethodPost, http.MethodOptions)

	router.HandleFunc("/api/user/{id}", controller.UserFromIdHandler).Methods(http.MethodGet)
	router.HandleFunc("/api/user/profile/{publicId}", controller.UserFromPublicIdHandler).Methods(http.MethodGet)

	router.HandleFunc("/api/course/{id}", controller.CourseFromIdHandler).Methods(http.MethodGet)
	router.HandleFunc("/api/courses/{ids}", controller.CoursesFromIdHandler).Methods(http.MethodGet)
	router.HandleFunc("/api/course/create", controller.CreateCourseHandler).Methods(http.MethodPost, http.MethodOptions)

	log.Fatal(http.ListenAndServe(":8080", router))
}
