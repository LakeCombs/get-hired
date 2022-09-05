package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID              primitive.ObjectID `bson:"_id"`
	Title           *string            `bson:"title" validate:"required"`
	Content         *string            `bson:"content" validate:"required"`
	Organization_id primitive.ObjectID `bson:"organization_id"  validate:"required"`
	Task_manager    primitive.ObjectID `bson:"task_manager,omitempty"`
	Created_at      time.Time          `bson:"created_at"`
	Updated_at      time.Time          `bson:"updated_at"`
	Task_status     *string            `bson:"task_status" validate:"required,eq=created|completed|in_progress|canceled|on_hold"`
}
