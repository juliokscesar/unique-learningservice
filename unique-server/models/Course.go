package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`

	Title       string `bson:"title" json:"title"`
	Subtitle	string `bson:"subtitle" json:"subtitle"`
	Description string `bson:"description" json:"description"`

	Teachers []primitive.ObjectID `bson:"teachers_id" json:"teachers_id"`
	Students []primitive.ObjectID `bson:"students_id" json:"students_id"`
}

func NewCourse(title, subtitle, description string) *Course {
	c := &Course{
		ID: primitive.NewObjectID(),

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

		Title:       	   title,
		Subtitle: 		subtitle,
		Description: description,
	}

	return c
}

func (c *Course) updateDate() {
	c.UpdatedAt = time.Now()
}

func (c *Course) ChangeTitle(title string) {
	c.Title = title
	c.updateDate()
}

func (c *Course) ChangeSubtitle(subtitle string) {
	c.Subtitle = subtitle
	c.updateDate()
}

func (c *Course) ChangeDescription(desc string) {
	c.Description = desc
	c.updateDate()
}

func (c *Course) AppendTeachers(teachers ...primitive.ObjectID) {
	c.Teachers = append(c.Teachers, teachers...)
	c.updateDate()
}

func (c *Course) AppendStudents(students ...primitive.ObjectID) {
	c.Students = append(c.Students, students...)
	c.updateDate()
}
