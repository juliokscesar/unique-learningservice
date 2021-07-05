package controller

import (
	"github.com/juliokscesar/unique-learningservice/unique-server/models"
)

func GetMaterialFromId(id string) (*models.Material, error) {
	m := new(models.Material)
	err := getById(MATERIALS_COLLECTION, id, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func InsertOneMaterial(m *models.Material) error {
	return insertOne(MATERIALS_COLLECTION, m)
}

func DeleteOneMaterial(id string) error {
	return deleteById(MATERIALS_COLLECTION, id)
}
