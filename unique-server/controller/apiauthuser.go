package controller

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/juliokscesar/unique-learningservice/unique-server/models"
	"github.com/juliokscesar/unique-learningservice/unique-server/uniqueErrors"
	"github.com/juliokscesar/unique-learningservice/unique-server/utils"
)

type ApiAuthUserController struct {}

func (auc *ApiAuthUserController) RegisterApiAuthUser(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r)

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

func CreateApiAuthUser(username, password string) (*models.ApiAuthUser, error) {
	_, err := GetApiAuthUserByUsername(username)
	if err == nil {
		return nil, uniqueErrors.ErrAPIUsernameRegistered
	}
	
	au, err := models.NewApiAuthUser(username, password)
	if err != nil {
		return nil, err
	}

	err = insertOne(API_AUTH_USERS_COLLECTION, au)
	if err != nil {
		return nil, err
	}

	return au, nil
}

func GetApiAuthUserByUsername(username string) (*models.ApiAuthUser, error) {
	au := new(models.ApiAuthUser)
	filter := bson.D{primitive.E{Key: "username", Value: username}}
	err := getByFilter(API_AUTH_USERS_COLLECTION, filter, au)
	if err != nil {
		return nil, err
	}

	return au, nil
}

func CheckApiAuthUser(username, password string) error {
	au, err := GetApiAuthUserByUsername(username)
	if err != nil {
		return err
	}

	if !au.CheckPassword(password) {
		return uniqueErrors.ErrInvalidAPIPassword
	}

	return nil
}
