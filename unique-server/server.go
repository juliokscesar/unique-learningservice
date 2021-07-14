package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/juliokscesar/unique-learningservice/unique-server/controller"
)

func configHandlers(r *mux.Router) {
	r.HandleFunc("/api/authuser/register", controller.RegisterApiAuthUser).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/api/user/register", 
		controller.ProvideHandler(controller.RegisterUserHandler),
	).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/api/user/login", 
		controller.ProvideHandler(controller.LoginUserHandler),
	).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/api/user/{id}", 
		controller.ProvideHandler(controller.UserFromIdHandler),
	).Methods(http.MethodGet)
	r.HandleFunc("/api/user/profile/{publicId}", 
		controller.ProvideHandler(controller.UserFromPublicIdHandler),
	).Methods(http.MethodGet)
	r.HandleFunc("/api/user/{id}/settings/change/{field}", 
		controller.ProvideHandler(controller.ChangeUserField),
	).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/api/course/{id}", 
		controller.ProvideHandler(controller.CourseFromIdHandler),
	).Methods(http.MethodGet)
	r.HandleFunc("/api/courses/{ids}", 
		controller.ProvideHandler(controller.CoursesFromIdHandler),
	).Methods(http.MethodGet)
	r.HandleFunc("/api/course/create", 
		controller.ProvideHandler(controller.CreateCourseHandler),
	).Methods(http.MethodPost, http.MethodOptions)
}

func main() {
	err := controller.ControllerInit()
	if err != nil {
		log.Fatal("Controller initalizing errror:", err)
	}

	router := mux.NewRouter()
	configHandlers(router)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
