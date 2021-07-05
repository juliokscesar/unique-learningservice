package controller

import (
	"github.com/juliokscesar/unique-learningservice/unique-server/models"
)

func GetAssignmentFromId(id string) (*models.Assignment, error) {
	a := new(models.Assignment)
	err := getById(ASSIGNMENTS_COLLECTION, id, a)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func InsertOneAssignment(a *models.Assignment) error {
	return insertOne(ASSIGNMENTS_COLLECTION, a)
}

func DeleteOneAssignment(id string) error {
	return deleteById(ASSIGNMENTS_COLLECTION, id)
}
