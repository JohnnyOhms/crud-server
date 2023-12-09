package Config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
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

	database := client.Database("UserInfo")
	collection := database.Collection("info")

	err = database.CreateCollection(ctx, "info")
	if err != nil {
		fmt.Println("Database already existed")
	}

	return collection, nil
}
