package crud

import (
	"context"
	"encoding/json"
	"net/http"
	"workshop/client"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// GetOne returns everything from a collection
func GetOne(collection string, ID string, w http.ResponseWriter) ([]byte, error) {
	IDobj, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	filter := bson.M{
		"_id": IDobj,
	}
	col := client.GetCollection(collection)
	dataRaw := bson.M{}
	err = col.FindOne(context.TODO(), filter).Decode(&dataRaw)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	data, _ := json.Marshal(dataRaw)
	return data, err

}
