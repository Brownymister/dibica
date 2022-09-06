package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertIntoCards(card CardData) {
	client := InitConnection()

	coll := client.Database(os.Getenv("MONGODB_DATABASE_NAME")).Collection(os.Getenv("MONGODB_COLLECTION_NAME"))

	fmt.Print(card)

	coll.InsertOne(
		context.TODO(),
		bson.D{
			{Key: "_id", Value: card.Id},
			{Key: "name", Value: card.Name},
			{Key: "template", Value: card.Template},
			{Key: "message", Value: card.Message},
			{Key: "cardlink", Value: card.CardLink},
		},
	)
}
