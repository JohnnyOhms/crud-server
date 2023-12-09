package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Info struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	UserId  string             `bson:"userid,omitempty"`
	Name    string             `bson:"name,omitempty"`
	Email   string             `bson:"email,omitempty"`
	Number  int16              `bson:"number,omitempty"`
	Address string             `bson:"address,omitempty"`
}
