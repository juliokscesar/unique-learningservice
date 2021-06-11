package controller

import (
	"context"
	"errors"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client    *mongo.Client
	clientErr error           = nil
	ctx       context.Context = context.TODO()

	usersCollection     *mongo.Collection
	coursesCollection   *mongo.Collection
	materialsCollection *mongo.Collection

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

	client, clientErr = mongo.Connect(ctx, clientOpt)
	if clientErr != nil {
		return clientErr
	}
	log.Println("Successfully connected client (func ControllerInit)")

	clientErr = client.Ping(ctx, nil)
	if clientErr != nil {
		return clientErr
	}
	log.Println("Successfully pinged client (func ControllerInit)")

	usersCollection = client.Database(Db).Collection("users")
	coursesCollection = client.Database(Db).Collection("courses")
	materialsCollection = client.Database(Db).Collection("materials")
	//assignmentsCollection = client.Database(Db).Collection("assignments")

	return nil
}
