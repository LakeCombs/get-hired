package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Experience struct {
	Company         *string
	Position        *string
	Work_experience *string
	// From *
	// To *

}

type User struct {
	Id              primitive.ObjectID `json:"id" bson:"_id"`
	First_name      string             `bson:"first_name, omitempty"`
	Last_name       string             `bson:"last_name, omitempty"`
	Skills          []string           `bson:"skills, omitempty"`
	Work_experience []Experience       `bson:"experience, omitempty"`
	Password        string             `bson:"password, omitempty"`
}
