package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type ApiAuthUser struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`

	Username string `bson:"username" json:"username"`
	PasswordHash string `bson:"password" json:"password"`
}

func NewApiAuthUser(username, password string) (*ApiAuthUser, error) {
	au := &ApiAuthUser{
		ID: primitive.NewObjectID(),

		CreatedAt: time.Now(),

		Username: username,
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), HashCost)
	if err != nil {
		return nil, err
	}

	au.PasswordHash = string(passHash)

	return au, nil
}

func (au *ApiAuthUser) CheckPassword(password string) bool {
	return (bcrypt.CompareHashAndPassword([]byte(au.PasswordHash), []byte(password)) == nil)
}
