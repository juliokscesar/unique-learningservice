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
	mongoClient    *mongo.Client
	mongoClientErr error
	ctx            context.Context = context.TODO()

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
	return mongoClient != nil
}

func ControllerInit() error {
	clientOpt := options.Client().ApplyURI(MongoURI)

	mongoClient, mongoClientErr = mongo.Connect(ctx, clientOpt)
	if mongoClientErr != nil {
		return mongoClientErr
	}
	log.Println("Successfully connected client (func ControllerInit) and client == nil?", mongoClient == nil)

	mongoClientErr = mongoClient.Ping(ctx, nil)
	if mongoClientErr != nil {
		return mongoClientErr
	}
	log.Println("Successfully pinged client (func ControllerInit)")

	usersCollection = mongoClient.Database(Db).Collection("users")
	coursesCollection = mongoClient.Database(Db).Collection("courses")
	materialsCollection = mongoClient.Database(Db).Collection("materials")
	assignmentsCollection = mongoClient.Database(Db).Collection("assignments")

	return nil
}

func ValidateConvertId(id string) (primitive.ObjectID, error) {
	if !primitive.IsValidObjectID(id) {
		return primitive.NilObjectID, ERR_INVALID_ID
	}

	oid, err := primitive.ObjectIDFromHex(id)

	return oid, err
}
