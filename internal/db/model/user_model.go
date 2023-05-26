package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	First_name string             `bson:"first_name,omitempty"`
	Last_name  string             `bson:"last_name,omitempty"`
	Email      string             `bson:"email, omitempty"`
	Password   string             `bson:"password, omitempty"`
}
