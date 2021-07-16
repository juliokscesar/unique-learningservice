package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/juliokscesar/unique-learningservice/unique-server/controller"
)

func configHandlers(r *mux.Router) {
	// User HTTP Handlers
	uc := new(controller.UserController)

	r.HandleFunc("/api/user/register",
		controller.ProvideHandler(uc.RegisterUser),
	).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/api/user/login",
		controller.ProvideHandler(uc.LoginUser),
	).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/api/user/{id}",
		controller.ProvideHandler(uc.UserFromId),
	).Methods(http.MethodGet)
	r.HandleFunc("/api/user/profile/{publicId}",
		controller.ProvideHandler(uc.UserFromPublicId),
	).Methods(http.MethodGet)
	r.HandleFunc("/api/user/{id}/settings/change/{field}",
		controller.ProvideHandler(uc.ChangeUserField),
	).Methods(http.MethodPost, http.MethodOptions)

	// Course HTTP Handlers
	cc := new(controller.CourseController)

	r.HandleFunc("/api/course/{id}",
		controller.ProvideHandler(cc.CourseFromId),
	).Methods(http.MethodGet)
	r.HandleFunc("/api/courses/{ids}",
		controller.ProvideHandler(cc.CoursesFromIds),
	).Methods(http.MethodGet)
	r.HandleFunc("/api/course/create",
		controller.ProvideHandler(cc.CreateCourse),
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
		AllowedHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
	})

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
