package controller

import (
	"github.com/juliokscesar/unique-learningservice/unique-server/models"
	"github.com/juliokscesar/unique-learningservice/unique-server/utils"
)

func CreateAndInsertCourse(title, subtitle, desc, admId string) (*models.Course, error) {
	c := models.NewCourse(title, subtitle, desc)

	convertedId, err := utils.ValidateConvertId(admId)
	if err != nil {
		return nil, err
	}

	c.AppendTeachers(convertedId)

	err = InsertOneCourse(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func GetCourseFromId(id string) (*models.Course, error) {
	c := new(models.Course)
	err := getById(COURSES_COLLECTION, id, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func GetManyCoursesFromId(ids []string) ([]*models.Course, error) {
	c := make([]*models.Course, len(ids))
	err := getManyById(COURSES_COLLECTION, ids, &c)
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
