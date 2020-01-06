package crud

import (
	"context"
	"net/http"
	"workshop/client"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Insert inserts stuff
func Insert(collection string, data interface{}, w http.ResponseWriter) (*mongo.InsertOneResult, error) {
	col := client.GetCollection(collection)
	result, err := col.InsertOne(context.TODO(), data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	return result, err

}

// GetAll returns everything from a collection
func GetAll(collection string, w http.ResponseWriter) (*mongo.Cursor, error) {
	col := client.GetCollection(collection)
	cursor, err := col.Find(context.TODO(), bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	return cursor, err

}
