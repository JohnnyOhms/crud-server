package model

import (
	"context"
	"fmt"

	"github.com/JohnnyOhms/crud-server/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Info struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserId      string             `bson:"userid,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Email       string             `bson:"email,omitempty"`
	Number      int                `bson:"number,omitempty"`
	Address     string             `bson:"address,omitempty"`
	DateCreated string
}

func InsertInfo(info Info, collection *mongo.Collection) (*mongo.InsertOneResult, error) {
	res, err := collection.InsertOne(context.Background(), info)
	if err != nil {
		fmt.Println("failed to add document to database")
		return nil, err
	}

	return res, nil
}

func GetInfo(userId string, collection *mongo.Collection) (primitive.M, error) {
	filter := bson.M{"userid": userId}
	var result bson.M

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println("failed to get document from database")
		return nil, err
	}

	cursor.Close((context.Background()))

	for cursor.Next(context.Background()) {

		if err := cursor.Decode(&result); err != nil {
			fmt.Println("failed to decode documents")
			return nil, err
		}
		// fmt.Println(result)
		if err := cursor.Err(); err != nil {
			fmt.Println("failed to loop throught all document")
			return nil, err
		}
	}
	return result, err
}

func GetSingleInfo(userId string, id interface{}, collection *mongo.Collection) (primitive.M, error) {
	filter := bson.M{
		"$and": []bson.M{
			{"userid": userId},
			{"_id": id},
		},
	}

	var result bson.M

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println("failed to get document from database")
		return nil, err
	}

	cursor.Close((context.Background()))

	for cursor.Next(context.Background()) {

		if err := cursor.Decode(&result); err != nil {
			fmt.Println("failed to decode documents")
			return nil, err
		}

		if err := cursor.Err(); err != nil {
			fmt.Println("failed to loop throught all document")
			return nil, err
		}
	}
	return result, err
}

func DeleteInfo(userId string, collection *mongo.Collection) (*mongo.DeleteResult, error) {
	filter := bson.M{"userid": userId}
	result, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		fmt.Println("failed to delete documents")
		return nil, err
	}
	return result, nil

}

func DeleteSingleInfo(userId string, id interface{}, collection *mongo.Collection) (*mongo.DeleteResult, error) {
	filter := bson.M{
		"$and": []bson.M{
			{"userid": userId},
			{"_id": id},
		},
	}
	result, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		fmt.Println("failed to delete documents")
		return nil, err
	}
	return result, nil
}

func EditInfo(userId string, id interface{}, data interface{}, collection *mongo.Collection) (*mongo.UpdateResult, error) {
	filter := bson.M{
		"$and": []bson.M{
			{"userid": userId},
			{"_id": id},
		},
	}

	bsonMap, err := utils.Encoding(data)
	if err != nil {
		fmt.Println("failed to encode data")
		return nil, err
	}

	var bsonMapConverted map[string]interface{}
	err = utils.Decoding(bsonMap, &bsonMapConverted)
	if err != nil {
		fmt.Println("failed to decode data")
		return nil, err
	}

	update := bson.M{"$set": bsonMapConverted}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("failed to update documents")
		return nil, err
	}
	return result, nil
}
