package bd

import (
	"context"
	"github.com/ivchip/login-go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func InsertRegister(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	usersCollection := MongoCN.Database(nameDB).Collection("users")
	u.Password, _ = EncryptPassword(u.Password)
	result, err := usersCollection.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}