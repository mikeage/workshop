package crud

import (
	"context"
	"encoding/json"
	"net/http"
	"workshop/client"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Insert will insert a document to the mongoDB
func Insert(collection string, data interface{}, w http.ResponseWriter) (*mongo.InsertOneResult, error) {

	col := client.GetCollection(collection)
	result, err := col.InsertOne(context.TODO(), data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	return result, err
}

//GetAll get all the documents
func GetAll(collection string, w http.ResponseWriter) (*mongo.Cursor, error) {
	col := client.GetCollection(collection)

	cursor, err := col.Find(context.TODO(), bson.M{})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	return cursor, err
}

//GetOne gets a single item from the database
func GetOne(collection, ID string, w http.ResponseWriter) ([]byte, error) {
	IDobj, err := primitive.ObjectIDFromHex(ID)
	col := client.GetCollection(collection)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return []byte{}, err
	}

	filter := bson.M{
		"_id": IDobj,
	}

	dataRaw := bson.M{}

	err = col.FindOne(context.TODO(), filter).Decode(&dataRaw)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return []byte{}, err
	}

	data, _ := json.Marshal(dataRaw)
	return data, nil

}

//Update will update an document in the mongoDB
func Update(collection, ID string, data interface{}, w http.ResponseWriter) (*mongo.UpdateResult, error) {
	IDobj, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return &mongo.UpdateResult{}, err
	}

	col := client.GetCollection(collection)

	filter := bson.M{
		"_id": IDobj,
	}

	opts := options.Replace().SetUpsert(true)

	result, err := col.ReplaceOne(context.TODO(), filter, data, opts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	return result, err
}
