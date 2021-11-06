package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Contact struct {
	ID		primitive.ObjectID	`json:"id,omitempty" bson:"_id,omitempty"`
	UserID	string	`json:"user_id,omitempty" bson:"user_id"`
	Name	string	`json:"name" bson:"name" binding:"required"`
	Email	string	`json:"email,omitempty" bson:"email,omitempty" binding:"required"`
	Phone	string	`json:"phone,omitempty" bson:"phone,omitempty" binding:"required"`
	Type	string	`json:"type,omitempty" bson:"type,omitempty"`
	Date	time.Time `json:"date,omitempty" bson:"date,omitempty"`
}
