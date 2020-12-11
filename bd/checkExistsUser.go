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

	db := MongoCN.Database("db_login")
	col := db.Collection("users")

	conditional := bson.M{"email": email}
	var result models.User
	err := col.FindOne(ctx, conditional).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}