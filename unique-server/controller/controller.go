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

func getManyByFilter(collection string, filter interface{}, decodeTo interface{}) error {
	r, err := collections[collection].Find(ctx, filter)
	if err != nil {
		return err
	}

	return r.All(ctx, decodeTo)
}

func getManyById(collection string, ids []string, decodeTo interface{}) error {
	idsLen := len(ids)
	convertedIds := make([]primitive.ObjectID, idsLen)

	for _, id := range ids {
		convertedId, err := utils.ValidateConvertId(id)
		if err != nil {
			return err
		}
		convertedIds = append(convertedIds, convertedId)
	}

	filter := bson.D{
		primitive.E{Key: "_id", Value: bson.D{
			primitive.E{Key: "$in", Value: convertedIds},
		}},
	}

	return getManyByFilter(collection, filter, decodeTo)
}

// Controller insert functions
func insertOne(collection string, obj interface{}) error {
	if !IsControllerInit() {
		return uniqueErrors.ErrNotInitialized
	}

	_, err := collections[collection].InsertOne(ctx, obj)
	return err
}

// Controller update functions
func updateOne(collection string, filter interface{}, updateOp interface{}) error {
	if !IsControllerInit() {
		return uniqueErrors.ErrNotInitialized
	}

	_, err := collections[collection].UpdateOne(ctx, filter, updateOp)
	if err != nil {
		return err
	}
	
	return nil
}

func updateById(collection string, id string, updateOp interface{}) error {
	convertedId, err := utils.ValidateConvertId(id)
	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: convertedId}}
	
	return updateOne(collection, filter, updateOp)
}

// Controller delete functions
func deleteOneByFilter(collection string, filter primitive.D) error {
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
	return deleteOneByFilter(collection, filter)
}
