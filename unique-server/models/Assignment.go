package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Assignment struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`

	Info            Material  `bson:"info" json:"info"`
	Deadline        time.Time `bson:"deadline" json:"deadline"`
	BlockOnDeadline bool      `bson:"block_on_deadline" json:"block_on_deadline"`
}

func NewAssignment(title string, description string, courseOrigin primitive.ObjectID, deadline time.Time, blockOnDeadline bool) *Assignment {
	mInfo := NewMaterial(title, description, courseOrigin)

	a := &Assignment{
		ID: primitive.NewObjectID(),

		Info:            *mInfo,
		Deadline:        deadline,
		BlockOnDeadline: blockOnDeadline,
	}

	return a
}

func NewAssignmentFromMaterial(info Material, deadline time.Time, blockOnDeadline bool) *Assignment {
	a := &Assignment{
		ID: primitive.NewObjectID(),

		Info:            info,
		Deadline:        deadline,
		BlockOnDeadline: blockOnDeadline,
	}

	return a
}

func (a *Assignment) updateDate() {
	a.Info.updateDate()
}

func (a *Assignment) ChangeDeadline(newDeadline time.Time) {
	a.Deadline = newDeadline

	a.Info.updateDate()
}

func (a *Assignment) SetBlock(b bool) {
	a.BlockOnDeadline = b

	a.Info.updateDate()
}
