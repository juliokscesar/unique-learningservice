package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Material struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`

	Title       string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`

	CourseOrigin primitive.ObjectID `bson:"course_origin" json:"course_origin"`
}

func NewMaterial(title string, description string, courseOrigin primitive.ObjectID) *Material {
	m := &Material{
		ID: primitive.NewObjectID(),

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

		Title:       title,
		Description: description,

		CourseOrigin: courseOrigin,
	}

	return m
}

func (m *Material) updateDate() {
	m.UpdatedAt = time.Now()
}

func (m *Material) ChangeTitle(newTitle string) {
	m.Title = newTitle
	m.updateDate()
}

func (m *Material) ChangeDescription(description string) {
	m.Description = description
	m.updateDate()
}
