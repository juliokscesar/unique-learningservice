package controller

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/juliokscesar/unique-learningservice/unique-server/models"
)

func GetCourseFromId(id string) (*models.Course, error) {
	if !IsControllerInit() {
		return nil, ERR_NOT_INITIALIZED
	}

	cid, err := ValidateConvertId(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: cid}}

	c := new(models.Course)
	err = coursesCollection.FindOne(ctx, filter).Decode(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func InsertOneCourse(c *models.Course) error {
	if !IsControllerInit() {
		return ERR_NOT_INITIALIZED
	}

	_, err := coursesCollection.InsertOne(ctx, c)
	if err != nil {
		return err
	}

	return nil
}

func DeleteOneCourse(id string) error {
	if !IsControllerInit() {
		return ERR_NOT_INITIALIZED
	}

	cid, err := ValidateConvertId(id)
	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: cid}}
	_, err = coursesCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
