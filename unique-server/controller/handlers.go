package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/juliokscesar/unique-learningservice/unique-server/models"
	"github.com/juliokscesar/unique-learningservice/unique-server/utils"
)

func errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	w.Write([]byte(`{"error": "` + err.Error() + `"}`))
}

func setupHandler(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.Header().Set("Content-Type", "application/json")
}

// User HTTP handlers
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	setupHandler(w, r)

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

	err = RegisterUser(u)
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(u)
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	setupHandler(w, r)

	err := r.ParseForm()
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	u, err := LoginUser(email, password)
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(u)
}

func UserFromIdHandler(w http.ResponseWriter, r *http.Request) {
	setupHandler(w, r)

	uid := mux.Vars(r)["id"]

	u, err := GetUserFromID(uid)
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(u)
}


func UserFromPublicIdHandler(w http.ResponseWriter, r *http.Request) {
	setupHandler(w, r)

	publicId := mux.Vars(r)["publicId"]

	u, err := GetUserFromPublicId(publicId)
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(u)
}

// Course HTTP handlers
func CreateCourseHandler(w http.ResponseWriter, r *http.Request) {
	setupHandler(w, r)
	
	err := r.ParseForm()
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	title := r.FormValue("title")
	subtitle := r.FormValue("subtitle")
	description := r.FormValue("description")
	admId := r.FormValue("admId")

	c, err := CreateAndInsertCourse(title, subtitle, description, admId)
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	err = AddUserCourse(admId, c.ID.Hex())
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(c)
}

func CourseFromIdHandler(w http.ResponseWriter, r *http.Request) {
	setupHandler(w, r)

	cid := mux.Vars(r)["id"]

	c, err := GetCourseFromId(cid)
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(c)
}

func CoursesFromIdHandler(w http.ResponseWriter, r *http.Request) {
	setupHandler(w, r)
	
	cids := mux.Vars(r)["ids"]
	
	coursesIds := strings.Split(cids, ",")

	courses, err := GetManyCoursesFromId(coursesIds)
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(courses)
}
