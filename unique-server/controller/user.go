package controller

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/juliokscesar/unique-learningservice/unique-server/models"
)

func GetUserFromID(id string) (*models.User, error) {
	if !IsControllerInit() {
		return nil, ERR_NOT_INITIALIZED
	}

	if !primitive.IsValidObjectID(id) {
		return nil, ERR_INVALID_ID
	}

	uid, err := primitive.ObjectIDFromHex(id)
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

func InsertOneUser(u *models.User) error {
	if !IsControllerInit() {
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

	if !primitive.IsValidObjectID(id) {
		return ERR_INVALID_ID
	}

	uid, err := primitive.ObjectIDFromHex(id)
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
