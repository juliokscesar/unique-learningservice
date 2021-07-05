package controller

import (
	"github.com/juliokscesar/unique-learningservice/unique-server/models"
)

func GetCourseFromId(id string) (*models.Course, error) {
	c := new(models.Course)
	err := getById(COURSES_COLLECTION, id, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func InsertOneCourse(c *models.Course) error {
	return insertOne(COURSES_COLLECTION, c)
}

func DeleteOneCourse(id string) error {
	return deleteById(COURSES_COLLECTION, id)
}
