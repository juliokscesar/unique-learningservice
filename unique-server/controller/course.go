package controller

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/juliokscesar/unique-learningservice/unique-server/models"
	"github.com/juliokscesar/unique-learningservice/unique-server/uniqueErrors"
	"github.com/juliokscesar/unique-learningservice/unique-server/utils"
)

func GetCourseFromId(id string) (*models.Course, error) {
	if !IsControllerInit() {
		return nil, uniqueErrors.ErrNotInitialized
	}

	cid, err := utils.ValidateConvertId(id)
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
		return uniqueErrors.ErrNotInitialized
	}

	_, err := coursesCollection.InsertOne(ctx, c)
	if err != nil {
		return err
	}

	return nil
}

func DeleteOneCourse(id string) error {
	if !IsControllerInit() {
		return uniqueErrors.ErrNotInitialized
	}

	cid, err := utils.ValidateConvertId(id)
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
