package bd

import (
	"context"
	"github.com/ivchip/login-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func CheckIsExitsUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	usersCollection := MongoCN.Database(nameDB).Collection("users")

	conditional := bson.M{"email": email}
	var result models.User
	err := usersCollection.FindOne(ctx, conditional).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}