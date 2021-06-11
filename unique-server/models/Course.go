package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`

	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`

	Teachers []primitive.ObjectID `bson:"teachers_id" json:"teachers_id"`
	Students []primitive.ObjectID `bson:"students_id" json:"students_id"`
}

func NewCourse(name, description string) *Course {
	c := &Course{
		ID: primitive.NewObjectID(),

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

		Name:        name,
		Description: description,
	}

	return c
}

func (c *Course) updateDate() {
	c.UpdatedAt = time.Now()
}

func (c *Course) AppendTeachers(teachers ...primitive.ObjectID) {
	c.Teachers = append(c.Teachers, teachers...)
	c.updateDate()
}

func (c *Course) AppendStudents(students ...primitive.ObjectID) {
	c.Students = append(c.Students, students...)
	c.updateDate()
}
