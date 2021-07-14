package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/juliokscesar/unique-learningservice/unique-server/models"
	"github.com/juliokscesar/unique-learningservice/unique-server/utils"
)

type CourseController struct {}

func (cc *CourseController) CreateCourse(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	title := r.FormValue("title")
	subtitle := r.FormValue("subtitle")
	description := r.FormValue("description")
	admId := r.FormValue("admId")

	c, err := createAndInsertCourse(title, subtitle, description, admId)
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	err = addUserCourse(admId, c.ID.Hex())
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(c)
}

func (cc *CourseController) CourseFromId(w http.ResponseWriter, r *http.Request) {
	cid := mux.Vars(r)["id"]

	c, err := getCourseFromId(cid)
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(c)
}

func (cc *CourseController) CoursesFromIds(w http.ResponseWriter, r *http.Request) {
	cids := mux.Vars(r)["ids"]

	coursesIds := strings.Split(cids, ",")

	courses, err := getManyCoursesFromId(coursesIds)
	if err != nil {
		errorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(courses)
}

func createAndInsertCourse(title, subtitle, desc, admId string) (*models.Course, error) {
	c := models.NewCourse(title, subtitle, desc)

	convertedId, err := utils.ValidateConvertId(admId)
	if err != nil {
		return nil, err
	}

	c.AppendTeachers(convertedId)

	err = insertOneCourse(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func getCourseFromId(id string) (*models.Course, error) {
	c := new(models.Course)
	err := getById(COURSES_COLLECTION, id, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func getManyCoursesFromId(ids []string) ([]*models.Course, error) {
	c := make([]*models.Course, len(ids))
	err := getManyById(COURSES_COLLECTION, ids, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func insertOneCourse(c *models.Course) error {
	return insertOne(COURSES_COLLECTION, c)
}

// func deleteOneCourse(id string) error {
// 	return deleteById(COURSES_COLLECTION, id)
// }
