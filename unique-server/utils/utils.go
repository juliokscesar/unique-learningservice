package utils

import (
	"log"
	"net/http"
	"net/mail"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/juliokscesar/unique-learningservice/unique-server/uniqueErrors"
)

func LogRequest(r *http.Request) {
	log.Println(r.Method, r.URL.Path, "by", r.RemoteAddr)
}

func ValidateConvertId(id string) (primitive.ObjectID, error) {
	if !primitive.IsValidObjectID(id) {
		return primitive.NilObjectID, uniqueErrors.ErrInvalidId
	}

	oid, err := primitive.ObjectIDFromHex(id)

	return oid, err
}

func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)

	if err != nil {
		return uniqueErrors.ErrInvalidEmail
	} else {
		return nil
	}
}
