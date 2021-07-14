package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/juliokscesar/unique-learningservice/unique-server/models"
	"github.com/juliokscesar/unique-learningservice/unique-server/uniqueErrors"
	"github.com/juliokscesar/unique-learningservice/unique-server/utils"
)

func errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	w.Write([]byte(`{"error": "` + err.Error() + `"}`))
}

func checkAuthentication(user, pass string) bool {
	return (CheckApiAuthUser(user, pass) == nil)
}

func setupHandler(w http.ResponseWriter, r *http.Request) error {
	user, pass, ok := r.BasicAuth()
	if !ok || !checkAuthentication(user, pass) {
		return uniqueErrors.ErrInvalidAPIAuthUser
	}

	w.Header().Set("Content-Type", "application/json")

	return nil
}

func ProvideHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)

		err := setupHandler(w, r)
		if err != nil {
			errorHandler(w, r, err)
			return
		}

		handler(w, r)
	}
}

// Api Auth User Handler (doesn't need authentication, so no setupHandler())
func RegisterApiAuthUser(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.Header().Set("Content-Type", "application/json")

	err := r.ParseForm()
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	au, err := CreateApiAuthUser(username, password)
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(au)
}

// User HTTP handlers
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
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
	uid := mux.Vars(r)["id"]

	u, err := GetUserFromID(uid)
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(u)
}

func UserFromPublicIdHandler(w http.ResponseWriter, r *http.Request) {
	publicId := mux.Vars(r)["publicId"]

	u, err := GetUserFromPublicId(publicId)
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(u)
}

func ChangeUserField(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	u := new(models.User)

	uid := mux.Vars(r)["id"]
	switch field := mux.Vars(r)["field"]; field {
	case "email":
		u, err = ChangeUserEmail(uid, r.FormValue("newEmail"))

	case "name":
		u, err = ChangeUserName(uid, r.FormValue("newName"))
	
	case "password":
		u, err = ChangeUserPass(uid, r.FormValue("oldPass"), r.FormValue("newPass"))

	default:
		errorHandler(w, r, uniqueErrors.ErrInvalidAPIUri)
		return
	}

	if err != nil {
		errorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(u)
}

// Course HTTP handlers
func CreateCourseHandler(w http.ResponseWriter, r *http.Request) {
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
	cid := mux.Vars(r)["id"]

	c, err := GetCourseFromId(cid)
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(c)
}

func CoursesFromIdHandler(w http.ResponseWriter, r *http.Request) {
	cids := mux.Vars(r)["ids"]

	coursesIds := strings.Split(cids, ",")

	courses, err := GetManyCoursesFromId(coursesIds)
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(courses)
}
