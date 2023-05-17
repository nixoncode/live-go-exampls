// Package database
// Created by GoLand.
// User: nixon
// Date: 17/5/2023
// Time: 12:47
package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var MongoClient *mongo.Client

func StartMongoDB() error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable")
	}
	var err error
	MongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return nil
}

func CloseMongoDB() {
	err := MongoClient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
}

func GetCollection(name string) *mongo.Collection {
	return MongoClient.Database("rest-api-go").Collection(name)
}
