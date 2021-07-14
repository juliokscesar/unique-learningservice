package controller

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"
	"github.com/juliokscesar/unique-learningservice/unique-server/models"
	"github.com/juliokscesar/unique-learningservice/unique-server/uniqueErrors"
	"github.com/juliokscesar/unique-learningservice/unique-server/utils"
)

type UserController struct {}

func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
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

	err = registerUser(u)
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(u)
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	u, err := loginUser(email, password)
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(u)
}

func (uc *UserController) UserFromId(w http.ResponseWriter, r *http.Request) {
	uid := mux.Vars(r)["id"]

	u, err := getUserFromID(uid)
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(u)
}

func (uc *UserController) UserFromPublicId(w http.ResponseWriter, r *http.Request) {
	publicId := mux.Vars(r)["publicId"]

	u, err := getUserFromPublicId(publicId)
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(u)
}

func (uc *UserController) ChangeUserField(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	u := new(models.User)

	uid := mux.Vars(r)["id"]
	switch field := mux.Vars(r)["field"]; field {
	case "email":
		u, err = changeUserEmail(uid, r.FormValue("newEmail"))

	case "name":
		u, err = changeUserName(uid, r.FormValue("newName"))
	
	case "password":
		u, err = changeUserPass(uid, r.FormValue("oldPass"), r.FormValue("newPass"))

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

func getUserFromID(id string) (*models.User, error) {
	u := new(models.User)
	err := getById(USERS_COLLECTION, id, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func getUserFromPublicId(publicId string) (*models.User, error) {
	u := new(models.User)
	
	filter := bson.D{primitive.E{Key: "public_id", Value: publicId}}
	err := getByFilter(USERS_COLLECTION, filter, u)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func getUserFromEmail(email string) (*models.User, error) {
	err := utils.ValidateEmail(email)
	if err != nil {
		return nil, err
	}

	filter := bson.D{primitive.E{Key: "email", Value: email}}

	u := new(models.User)
	err = getByFilter(USERS_COLLECTION, filter, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}


func loginUser(email, password string) (*models.User, error) {
	u, err := getUserFromEmail(email)
	if err != nil {
		return nil, uniqueErrors.ErrInvalidUserEmail
	}
	
	if !u.CheckPassword(password) {
		return nil, uniqueErrors.ErrInvalidPassword
	}
	
	return u, nil
}

func registerUser(u *models.User) error {
	_, err := getUserFromEmail(u.Email)
	if err == nil {
		return uniqueErrors.ErrEmailRegistered
	}
	
	return insertOneUser(u)
}

func addUserCourse(uid string, cid string) error {
	newU, err := getUserFromID(uid)
	if err != nil {
		return err
	}

	convertedCid, err := utils.ValidateConvertId(cid)
	if err != nil {
		return err
	}

	newU.AppendCourses(convertedCid)

	err = updateUserById(uid, bson.D{primitive.E{Key: "$set", Value: newU}})
	if err != nil {
		return err
	}

	return nil
}

func changeUserEmail(uid string, newEmail string) (*models.User, error) {
	err := utils.ValidateEmail(newEmail)
	if err != nil {
		return nil, err
	}

	u, err := getUserFromID(uid)
	if err != nil {
		return nil, err
	}

	u.UpdateEmail(newEmail)

	updateOp := bson.D{primitive.E{Key: "$set", Value: u}}
	err = updateUserById(uid, updateOp)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func changeUserName(uid string, newName string) (*models.User, error) {
	u, err := getUserFromID(uid)
	if err != nil {
		return nil, err
	}

	u.UpdateName(newName)

	updateOp := bson.D{primitive.E{Key: "$set", Value: u}}
	err = updateUserById(uid, updateOp)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func changeUserPass(uid string, oldPass string, newPass string) (*models.User, error) {
	u, err := getUserFromID(uid)
	if err != nil {
		return nil, err
	}

	if !u.CheckPassword(oldPass) {
		return nil, uniqueErrors.ErrInvalidPassword
	}

	err = u.UpdatePassword(newPass)
	if err != nil {
		return nil, err
	}

	updateOp := bson.D{primitive.E{Key: "$set", Value: u}}
	err = updateUserById(uid, updateOp)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func insertOneUser(u *models.User) error {
	return insertOne(USERS_COLLECTION, u)
}

func updateUserById(uid string, updateOp interface{}) error {
	return updateById(USERS_COLLECTION, uid, updateOp)
}

func DeleteOneUser(id string) error {
	// TODO: After courses controller, Delete user from course's students or teachers as well

	return deleteById(USERS_COLLECTION, id)
}
