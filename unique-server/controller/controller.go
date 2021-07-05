package controller

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/juliokscesar/unique-learningservice/unique-server/uniqueErrors"
	"github.com/juliokscesar/unique-learningservice/unique-server/utils"
)

var (
	mongoClient    *mongo.Client
	mongoClientErr error
	ctx            context.Context = context.TODO()

	collections map[string]*mongo.Collection = make(map[string]*mongo.Collection)

	MongoURI string = os.Getenv("MONGOURI_UNIQUE")
	Db       string = "unique_db"
)

const (
	USERS_COLLECTION       = "users"
	COURSES_COLLECTION     = "courses"
	MATERIALS_COLLECTION   = "materials"
	ASSIGNMENTS_COLLECTION = "assignments"
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
	log.Println("Successfully connected client (func ControllerInit)")

	mongoClientErr = mongoClient.Ping(ctx, nil)
	if mongoClientErr != nil {
		return mongoClientErr
	}
	log.Println("Successfully pinged client (func ControllerInit)")

	collections[USERS_COLLECTION] = mongoClient.Database(Db).Collection("users")
	collections[COURSES_COLLECTION] = mongoClient.Database(Db).Collection("courses")
	collections[MATERIALS_COLLECTION] = mongoClient.Database(Db).Collection("materials")
	collections[ASSIGNMENTS_COLLECTION] = mongoClient.Database(Db).Collection("assignments")

	return nil
}

// Controller Get functions

func getByFilter(collection string, filter primitive.D, decodeTo interface{}) error {
	if !IsControllerInit() {
		return uniqueErrors.ErrNotInitialized
	}

	err := collections[collection].FindOne(ctx, filter).Decode(decodeTo)
	return err
}

func getById(collection string, id string, decodeTo interface{}) error {
	convertedId, err := utils.ValidateConvertId(id)
	if err != nil {
		return err
	}

	filter := primitive.D{bson.E{Key: "_id", Value: convertedId}}

	err = getByFilter(collection, filter, decodeTo)
	return err
}

// Controller insert functions
func insertOne(collection string, obj interface{}) error {
	if !IsControllerInit() {
		return uniqueErrors.ErrNotInitialized
	}

	_, err := collections[collection].InsertOne(ctx, obj)
	return err
}

// Controller delete functions
func deleteOne(collection string, filter primitive.D) error {
	if !IsControllerInit() {
		return uniqueErrors.ErrNotInitialized
	}

	_, err := collections[collection].DeleteOne(ctx, filter)
	return err
}

func deleteById(collection string, id string) error {
	convertedId, err := utils.ValidateConvertId(id)
	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: convertedId}}
	return deleteOne(collection, filter)
}
