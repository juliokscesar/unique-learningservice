package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/juliokscesar/unique-learningservice/unique-server/utils"
)

const HashCost = 14

type User struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`

	PublicId string `bson:"public_id" json:"public_id"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`

	Name         string `bson:"name" json:"name"`
	Email        string `bson:"email" json:"email"`
	PasswordHash string `bson:"passwordHash" json:"passwordHash"`

	Courses []primitive.ObjectID `bson:"courses_id" json:"courses_id"`
}

func NewUser(name, email, password string) (*User, error) {
	u := &User{
		ID: primitive.NewObjectID(),
		
		PublicId: utils.RandSeq(24),

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

		Name:  name,
		Email: email,
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), HashCost)
	if err != nil {
		return nil, err
	}

	u.PasswordHash = string(passHash)

	return u, nil
}

func (u *User) AppendCourses(ids ...primitive.ObjectID) {
	u.Courses = append(u.Courses, ids...)
}

func (u *User) UpdateName(name string) {
	updateName(u, name)
}

func (u *User) UpdateEmail(email string) {
	updateEmail(u, email)
}

func (u *User) UpdatePassword(newPassword string) error {
	return updatePassword(u, newPassword)
}

func (u *User) CheckPassword(password string) bool {
	return checkPassword(u, password)
}

func (u *User) updateDate() {
	u.UpdatedAt = time.Now()
}

func updateName(u *User, name string) {
	u.Name = name
	u.updateDate()
}

func updateEmail(u *User, email string) {
	u.Email = email
	u.updateDate()
}

func updatePassword(u *User, newPassword string) error {
	newHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), HashCost)
	if err != nil {
		return err
	}

	u.PasswordHash = string(newHash)

	u.updateDate()
	return nil
}

func checkPassword(u *User, password string) bool {
	return (bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)) == nil)
}
