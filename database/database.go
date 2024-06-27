package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDatabase(docName string) *mongo.Collection  {
	// check if there are any error, if yes, it will return not nil, if it not, it will return nil
	err := godotenv.Load(".env") 
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	MONGO_URI := os.Getenv("MONGO_URI")

	clientOption := options.Client().ApplyURI(MONGO_URI)
	client, err = mongo.Connect(context.Background(), clientOption)
	if err != nil {
		log.Fatalf("Error while connect to MongoDB: %s", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error while ping to MongoDB: %s", err)
	}

	collection := client.Database("ClickNServe").Collection(docName)

	return collection
}

func DisconnectDatabase() {
	if client == nil {
		log.Fatalf("MongoDB is not initialized")
	}

	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatalf("Error while disconnecting from MongoDB: %s", err)
	}

	log.Println("Successfully disconnect from MongoDB")
}