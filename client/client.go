package client

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func initClient() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
}

//GetCollection gets the collection from the database
func GetCollection(collection string) *mongo.Collection {
	if client == nil {
		initClient()
	}

	ucol := client.Database("loanCalc").Collection(collection)

	return ucol
}
