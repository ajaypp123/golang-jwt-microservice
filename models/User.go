package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserId    string             `json:"userid"`
	FirstName *string            `json:"first_name"`
	LastName  *string            `json:"last_name"`
	Password  *string            `json:"password"`
	Email     *string            `json:"email"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"update_at"`
}

type UserDetail struct {
	User         *User   `json:"user"`
	Token        *string `json:"token"`
	RefreshToken *string `json:"refresh_token"`
}
