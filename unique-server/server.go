package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/juliokscesar/unique-learningservice/unique-server/controller"
	"github.com/juliokscesar/unique-learningservice/unique-server/models"
	"github.com/juliokscesar/unique-learningservice/unique-server/utils"
)

func errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	w.Write([]byte(`{"error": "` + err.Error() + `"}`))
	log.Println("Error:", err)
}

func registerUserHandler(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.Header().Set("Content-Type", "application/json")

	err := r.ParseForm()
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	u, err := models.NewUser(name, email, password)
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	err = controller.RegisterUser(u)
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(u)
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

	router.HandleFunc("/api/user/register", registerUserHandler).Methods(http.MethodPost, http.MethodOptions)

	log.Fatal(http.ListenAndServe(":8080", router))
}
