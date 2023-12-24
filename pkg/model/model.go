package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Info struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserId      string             `bson:"userid,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Email       string             `bson:"email,omitempty"`
	Phone       int                `bson:"phone,omitempty"`
	Address     string             `bson:"address,omitempty"`
	DateCreated string             `bson:"datecreated,omitempty"`
}

func InsertInfo(info Info, collection *mongo.Collection) (*mongo.InsertOneResult, error) {
	res, err := collection.InsertOne(context.Background(), info)
	if err != nil {
		fmt.Println("failed to add document to database")
		return nil, err
	}

	return res, nil
}

func GetInfo(userId string, collection *mongo.Collection) ([]primitive.M, error) {
	filter := bson.M{"userid": userId}
	var result []primitive.M

	ctx := context.Background()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to get document from database: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var eachResult primitive.M
		if err := cursor.Decode(&eachResult); err != nil {
			fmt.Println("Failed to decode document:", err)
			continue
		}
		result = append(result, eachResult)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate through documents: %w", err)
	}

	if len(result) == 0 {
		fmt.Println("No documents found for userId:", userId)
	}

	return result, nil
}

func GetSingleInfo(userId string, id string, collection *mongo.Collection) (primitive.M, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Failed to convert ID to ObjectID:", err)
		return nil, err
	}

	filter := bson.M{
		"$and": []bson.M{
			{"userid": userId},
			{"_id": objectID},
		},
	}

	var result primitive.M

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Failed to get document from database:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	if cursor.Next(context.Background()) {
		if err := cursor.Decode(&result); err != nil {
			fmt.Println("Failed to decode document:", err)
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("document not found")
	}

	return result, nil
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

func DeleteSingleInfo(userId string, id string, collection *mongo.Collection) (*mongo.DeleteResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Failed to convert ID to ObjectID:", err)
		return nil, err
	}

	filter := bson.M{
		"$and": []bson.M{
			{"userid": userId},
			{"_id": objectID},
		},
	}
	result, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		fmt.Println("failed to delete documents")
		return nil, err
	}
	return result, nil
}

func EditInfo(userId string, id string, data interface{}, collection *mongo.Collection) (*mongo.UpdateResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Failed to convert ID to ObjectID:", err)
		return nil, err
	}

	filter := bson.M{
		"userid": userId,
		"_id":    objectID,
	}

	updateData, err := bson.Marshal(data)
	if err != nil {
		fmt.Println("failed to marshal data:", err)
		return nil, err
	}

	var updateDoc bson.M
	err = bson.Unmarshal(updateData, &updateDoc)
	if err != nil {
		fmt.Println("failed to unmarshal data:", err)
		return nil, err
	}

	update := bson.M{"$set": updateDoc}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("failed to update documents:", err)
		return nil, err
	}
	return result, nil
}

func DeleteAllUserAndInfo(collection *mongo.Collection) (*mongo.DeleteResult, error) {
	filter := bson.M{"email": bson.M{"$regex": primitive.Regex{Pattern: "@", Options: ""}}}
	result, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		fmt.Println("failed to delete documents")
		return nil, err
	}
	return result, nil

}
