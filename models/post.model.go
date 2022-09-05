package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID         primitive.ObjectID `bson:"_id"`
	User_id    primitive.ObjectID `bson:"user_id"`
	Post_id    primitive.ObjectID `bson:"post_id"`
	Created_at time.Time          `bson:"created_at"`
	Updated_at time.Time          `bson:"updated_at"`
}

type Post struct {
	ID         primitive.ObjectID   `bson:"_id"`
	Content    *string              `bson:"content" validate:"required"`
	Media      []string             `bson:"media,omitempty"`
	Created_at time.Time            `bson:"created_at"`
	Updated_at time.Time            `bson:"updated_at"`
	Likes      []primitive.ObjectID `bson:"likes,omitempty"`
	Love       []primitive.ObjectID `bson:"love, omitempty"`
	Comments   []primitive.ObjectID `bson:"comments,omitempty"`
	Shares     []primitive.ObjectID `bson:"shares,omitempty"`
	Tags       []string             `bson:"tags,omitempty"`
}
