package controller

import (
	"context"
	"log"
	"os"

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

	usersCollection = mongoClient.Database(Db).Collection("users")
	coursesCollection = mongoClient.Database(Db).Collection("courses")
	materialsCollection = mongoClient.Database(Db).Collection("materials")
	assignmentsCollection = mongoClient.Database(Db).Collection("assignments")

	return nil
}
