package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	UserId   string `bson:"userid,omitempty"`
	Username string `bson:"username,omitempty"`
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}

type LoginUser struct {
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}

func CreateUser(user User, collection *mongo.Collection) error {
	filter := bson.M{"email": user.Email}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("failed to make qurey in searching user", err)
		return err
	}
	defer cursor.Close(context.Background())

	if cursor.Next(context.Background()) {
		return fmt.Errorf("user exist in the database, try another email")
	} else {
		_, err := collection.InsertOne(context.Background(), user)
		if err != nil {
			fmt.Println("Failed create user, try again", err)
			return err
		}
	}
	return nil
}

func GetUser(user LoginUser, collection *mongo.Collection) (primitive.M, error) {
	var result primitive.M

	filter := bson.M{
		"$and": []bson.M{
			{"email": user.Email},
			{"password": user.Password},
		},
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Failed to get user from database:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	if cursor.Next(context.Background()) {
		if err := cursor.Decode(&result); err != nil {
			fmt.Println("Failed to decode document:", err)
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("user not found in the database")
	}

	return result, nil
}
