package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Link struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ShortUrl    string             `json:"short_url" bson:"short_url,omitempty"`
	OriginalUrl string             `json:"original_url" bson:"original_url,omitempty"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	Clicks      int                `json:"clicks" bson:"clicks,omitempty"`
}
