package types

import "go.mongodb.org/mongo-driver/bson/primitive"

// User something
type User struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
}
