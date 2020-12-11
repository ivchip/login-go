package bd

import (
	"context"
	"github.com/ivchip/login-go/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var userDB, _ = utils.GetValueEnvironment("DB.USERNAME")
var passwordDB, _ = utils.GetValueEnvironment("DB.PASSWORD")
var nameDB, _ = utils.GetValueEnvironment("DB.NAME")
var atlasUri = "mongodb+srv://"+ userDB + ":" + passwordDB + "@cluster0.oo5xp.mongodb.net/" + nameDB + "?retryWrites=true&w=majority"
var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI(atlasUri)

func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Connection success with DB")
	return client
}

func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}