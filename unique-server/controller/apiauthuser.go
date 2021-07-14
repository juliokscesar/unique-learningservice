package controller

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/juliokscesar/unique-learningservice/unique-server/models"
	"github.com/juliokscesar/unique-learningservice/unique-server/uniqueErrors"
)

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
