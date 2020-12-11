package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	FirstName string `bson:"firstName" json:"firstName,omitempty"`
	LastName string `bson:"lastName" json:"lastName,omitempty"`
	BirthDate time.Time `bson:"birthDate" json:"birthDate,omitempty"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password,omitempty"`
}