package controller

import (
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/juliokscesar/unique-learningservice/unique-server/models"
)

var (
	ERR_EMAIL_REGISTERED = errors.New("Email is already registered")
)

func GetUserFromID(id string) (*models.User, error) {
	if !IsControllerInit() {
		return nil, ERR_NOT_INITIALIZED
	}

	uid, err := ValidateConvertId(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: uid}}

	u := new(models.User)
	err = usersCollection.FindOne(ctx, filter).Decode(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func GetUserFromEmail(email string) (*models.User, error) {
	if !IsControllerInit() {
		fmt.Println("We are in GetUserFromEmail: IsControllerInit returned false.")
		fmt.Println("client == nil", mongoClient == nil)
		return nil, ERR_NOT_INITIALIZED
	}

	filter := bson.D{primitive.E{Key: "email", Value: email}}

	u := new(models.User)
	err := usersCollection.FindOne(ctx, filter).Decode(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func RegisterUser(u *models.User) error {
	_, err := GetUserFromEmail(u.Email)
	if err == nil {
		return ERR_EMAIL_REGISTERED
	}

	return insertOneUser(u)
}

func insertOneUser(u *models.User) error {
	if !IsControllerInit() {
		fmt.Println("we are in insertoneuser: iscontrollerinit failed")
		fmt.Println("client == nil,", mongoClient == nil)
		return ERR_NOT_INITIALIZED
	}

	_, err := usersCollection.InsertOne(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func DeleteOneUser(id string) error {
	// TODO: After courses controller, Delete user from course's students or teachers as well

	if !IsControllerInit() {
		return ERR_NOT_INITIALIZED
	}

	uid, err := ValidateConvertId(id)
	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: uid}}

	_, err = usersCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
