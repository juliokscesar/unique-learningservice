package controller

import (
	"context"
	"errors"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	ctx    context.Context = context.TODO()

	usersCollection       *mongo.Collection
	coursesCollection     *mongo.Collection
	materialsCollection   *mongo.Collection
	assignmentsCollection *mongo.Collection

	MongoURI string = os.Getenv("MONGOURI_UNIQUE")
	Db       string = "unique_db"

	// Errors
	ERR_NOT_INITIALIZED = errors.New("Controller not initialized")
	ERR_INVALID_ID      = errors.New("Invalid Object ID")
)

func IsControllerInit() bool {
	return client != nil
}

func ControllerInit() error {
	clientOpt := options.Client().ApplyURI(MongoURI)

	client, err := mongo.Connect(ctx, clientOpt)
	if err != nil {
		return err
	}
	log.Println("Successfully connected client (func ControllerInit)")

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}
	log.Println("Successfully pinged client (func ControllerInit)")

	usersCollection = client.Database(Db).Collection("users")
	coursesCollection = client.Database(Db).Collection("courses")
	materialsCollection = client.Database(Db).Collection("materials")
	assignmentsCollection = client.Database(Db).Collection("assignments")

	return nil
}

func ValidateConvertId(id string) (primitive.ObjectID, error) {
	if !primitive.IsValidObjectID(id) {
		return primitive.NilObjectID, ERR_INVALID_ID
	}

	oid, err := primitive.ObjectIDFromHex(id)

	return oid, err
}
