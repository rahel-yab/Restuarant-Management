package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	First_name     *string            `bson:"first_name" json:"first_name" validate:"required,min=2,max=100"`
	Last_name      *string            `bson:"last_name" json:"last_name" validate:"required,min=2,max=100"`
	Profile_picture *string           `bson:"profile_picture,omitempty" json:"profile_picture,omitempty"`
	Is_verified    *bool              `bson:"is_verified" json:"is_verified"`
	Password       string             `bson:"password" json:"-"`
	Email          string             `bson:"email" json:"email" validate:"required,email"`
	Created_at     time.Time          `bson:"created_at" json:"created_at"`
	Updated_at     time.Time          `bson:"updated_at" json:"updated_at"`
	Access_token   string             `bson:"access_token,omitempty" json:"access_token,omitempty"`
	Refresh_token  string             `bson:"refresh_token,omitempty" json:"refresh_token,omitempty"`
	User_id        string             `bson:"user_id" json:"user_id"`
}