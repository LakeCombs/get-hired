package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Organization struct {
	Name     string               `json:"name" validate:"required"`
	Industry string               `json:"industry" validate:"required"`
	Staffs   []primitive.ObjectID `json:"staff,omitempty"`
	Interns  []primitive.ObjectID `json:"interns,omitempty"`
}
