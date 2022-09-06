package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

func GetCardById(id string) CardData {
	client := InitConnection()

	coll := client.Database(os.Getenv("MONGODB_DATABASE_NAME")).Collection(os.Getenv("MONGODB_COLLECTION_NAME"))

	var result CardData
	err := coll.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&result)
	if err != nil {
		panic(err)
	}

	return result
}
