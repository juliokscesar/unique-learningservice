package controller

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/juliokscesar/unique-learningservice/unique-server/models"
	"github.com/juliokscesar/unique-learningservice/unique-server/uniqueErrors"
	"github.com/juliokscesar/unique-learningservice/unique-server/utils"
)

func GetUserFromID(id string) (*models.User, error) {
	u := new(models.User)
	err := getById(USERS_COLLECTION, id, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func GetUserFromPublicId(publicId string) (*models.User, error) {
	u := new(models.User)
	
	filter := bson.D{primitive.E{Key: "public_id", Value: publicId}}
	err := getByFilter(USERS_COLLECTION, filter, u)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func GetUserFromEmail(email string) (*models.User, error) {
	err := utils.ValidateEmail(email)
	if err != nil {
		return nil, err
	}

	filter := bson.D{primitive.E{Key: "email", Value: email}}

	u := new(models.User)
	err = getByFilter(USERS_COLLECTION, filter, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}


func LoginUser(email, password string) (*models.User, error) {
	u, err := GetUserFromEmail(email)
	if err != nil {
		return nil, uniqueErrors.ErrInvalidUserEmail
	}
	
	if !u.CheckPassword(password) {
		return nil, uniqueErrors.ErrInvalidPassword
	}
	
	return u, nil
}

func RegisterUser(u *models.User) error {
	_, err := GetUserFromEmail(u.Email)
	if err == nil {
		return uniqueErrors.ErrEmailRegistered
	}
	
	return insertOneUser(u)
}

func AddUserCourse(uid string, cid string) error {
	newU, err := GetUserFromID(uid)
	if err != nil {
		return err
	}

	convertedCid, err := utils.ValidateConvertId(cid)
	if err != nil {
		return err
	}

	newU.AppendCourses(convertedCid)

	err = updateUserById(uid, bson.D{primitive.E{Key: "$set", Value: newU}})
	if err != nil {
		return err
	}

	return nil
}

func ChangeUserEmail(uid string, newEmail string) (*models.User, error) {
	err := utils.ValidateEmail(newEmail)
	if err != nil {
		return nil, err
	}

	u, err := GetUserFromID(uid)
	if err != nil {
		return nil, err
	}

	u.UpdateEmail(newEmail)

	updateOp := bson.D{primitive.E{Key: "$set", Value: u}}
	err = updateUserById(uid, updateOp)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func ChangeUserName(uid string, newName string) (*models.User, error) {
	u, err := GetUserFromID(uid)
	if err != nil {
		return nil, err
	}

	u.UpdateName(newName)

	updateOp := bson.D{primitive.E{Key: "$set", Value: u}}
	err = updateUserById(uid, updateOp)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func ChangeUserPass(uid string, oldPass string, newPass string) (*models.User, error) {
	u, err := GetUserFromID(uid)
	if err != nil {
		return nil, err
	}

	if !u.CheckPassword(oldPass) {
		return nil, uniqueErrors.ErrInvalidPassword
	}

	err = u.UpdatePassword(newPass)
	if err != nil {
		return nil, err
	}

	updateOp := bson.D{primitive.E{Key: "$set", Value: u}}
	err = updateUserById(uid, updateOp)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func insertOneUser(u *models.User) error {
	return insertOne(USERS_COLLECTION, u)
}

func updateUserById(uid string, updateOp interface{}) error {
	return updateById(USERS_COLLECTION, uid, updateOp)
}

func DeleteOneUser(id string) error {
	// TODO: After courses controller, Delete user from course's students or teachers as well

	return deleteById(USERS_COLLECTION, id)
}
