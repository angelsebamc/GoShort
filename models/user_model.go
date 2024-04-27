package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
	Email    string             `json:"email" bson:"email,omitempty"`
	Created  time.Time          `json:"created" bson:"created,omitempty" default:"time.Now()"`
}
