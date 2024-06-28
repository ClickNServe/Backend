package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func GetDatabase() *mongo.Database {
	return Client.Database("ClickNServe")
}

func ConnectDatabase()  {
	// check if there are any error, if yes, it will return not nil, if it not, it will return nil
	err := godotenv.Load(".env") 
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	MONGO_URI := os.Getenv("MONGO_URI")

	clientOption := options.Client().ApplyURI(MONGO_URI)
	Client, err = mongo.Connect(context.Background(), clientOption)
	if err != nil {
		log.Fatalf("Error while connect to MongoDB: %s", err)
	}

	err = Client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error while ping to MongoDB: %s", err)
	}
}

func DisconnectDatabase() {
	if Client == nil {
		log.Fatalf("MongoDB is not initialized")
	}

	err := Client.Disconnect(context.Background())
	if err != nil {
		log.Fatalf("Error while disconnecting from MongoDB: %s", err)
	}

	log.Println("Successfully disconnect from MongoDB")
}