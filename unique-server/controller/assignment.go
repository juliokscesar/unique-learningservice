package controller

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/juliokscesar/unique-learningservice/unique-server/models"
	"github.com/juliokscesar/unique-learningservice/unique-server/uniqueErrors"
	"github.com/juliokscesar/unique-learningservice/unique-server/utils"
)

func GetAssignmentFromId(id string) (*models.Assignment, error) {
	if !IsControllerInit() {
		return nil, uniqueErrors.ErrNotInitialized
	}

	aid, err := utils.ValidateConvertId(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: aid}}

	a := new(models.Assignment)
	err = assignmentsCollection.FindOne(ctx, filter).Decode(a)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func InsertOneAssignment(a *models.Assignment) error {
	if !IsControllerInit() {
		return uniqueErrors.ErrNotInitialized
	}

	_, err := assignmentsCollection.InsertOne(ctx, a)
	if err != nil {
		return err
	}

	return nil
}

func DeleteOneAssignment(id string) error {
	if !IsControllerInit() {
		return uniqueErrors.ErrNotInitialized
	}

	aid, err := utils.ValidateConvertId(id)
	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: aid}}
	_, err = assignmentsCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
