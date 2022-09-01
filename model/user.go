package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Experience struct {
	Company        *string
	Position       *string
	WorkExperience *string
	From           *string
	To             *string
}

type User struct {
	Id             primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	FirstName      string             `bson:"first_name, omitempty" json:"first_name,omitempty"`
	LastName       string             `bson:"last_name, omitempty" json:"last_name,omitempty"`
	Skills         []string           `bson:"skills, omitempty" json:"skills,omitempty"`
	WorkExperience []Experience       `bson:"experience, omitempty" json:"work_experience,omitempty"`
	Password       string             `bson:"password, omitempty" json:"password,omitempty"`
	Email          string             `bson:"email" json:"email,omitempty"`
	Location       string             `bson:"location" json:"location, omitempty"`
}
