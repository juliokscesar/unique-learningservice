package controller

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/juliokscesar/unique-learningservice/unique-server/models"
)

func GetMaterialFromId(id string) (*models.Material, error) {
	if !IsControllerInit() {
		return nil, ErrNotInitialized
	}

	mid, err := ValidateConvertId(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: mid}}

	m := new(models.Material)
	err = materialsCollection.FindOne(ctx, filter).Decode(m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func InsertOneMaterial(m *models.Material) error {
	if !IsControllerInit() {
		return ErrNotInitialized
	}

	_, err := materialsCollection.InsertOne(ctx, m)
	if err != nil {
		return err
	}

	return nil
}

func DeleteOneMaterial(id string) error {
	if !IsControllerInit() {
		return ErrNotInitialized
	}

	mid, err := ValidateConvertId(id)
	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: mid}}
	_, err = materialsCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
