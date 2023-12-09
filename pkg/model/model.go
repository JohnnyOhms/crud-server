package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Info struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	UserId  string             `bson:"userid,omitempty"`
	Name    string             `bson:"name,omitempty"`
	Email   string             `bson:"email,omitempty"`
	Number  int                `bson:"number,omitempty"`
	Address string             `bson:"address,omitempty"`
}

func InsertInfo(info Info, collection *mongo.Collection) (*mongo.InsertOneResult, error) {
	res, err := collection.InsertOne(context.Background(), info)
	if err != nil {
		fmt.Println("failed to add document to database")
		return nil, err
	}

	return res, nil
}
