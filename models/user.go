package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username     string             `json:"username" bson:"username"`
	PasswordHash string             `json:"passwordHash" bson:"passwordHash"`
	Email        string             `json:"email" bson:"email"`
	Role         string             `json:"role" bson:"role"`
	RefreshToken string             `json:"refreshToken" bson:"refreshToken"`
	CreatedAt    primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt    primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
}

