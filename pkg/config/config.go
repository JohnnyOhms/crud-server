package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection
var AuthCollection *mongo.Collection
var ctx = context.TODO()

func ConnectDB() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://officialjohn662:JohnnyOhms@cluster0.khdqojb.mongodb.net/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	Collection = client.Database("UserInfo").Collection("info")
	return Collection, nil
}

func ConnectAuth() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://officialjohn662:JohnnyOhms@cluster0.khdqojb.mongodb.net/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	AuthCollection = client.Database("UserInfo").Collection("user")
	return AuthCollection, nil
}
