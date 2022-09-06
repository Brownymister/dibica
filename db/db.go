package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CardData struct {
	Id       string `json:"id"`       // card id
	Name     string `json:"name"`     // name of the recever of the card
	Message  string `json:"message"`  // message on the backsite
	Template string `json:"template"` // card template
	CardLink string `json:"cardlink"` // image link
}

func InitConnection() *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().
		ApplyURI(os.Getenv("MONGODB_URI")).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, clientOptions)
	return client
}
