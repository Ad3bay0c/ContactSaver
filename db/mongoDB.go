package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

type MongoDB struct {
	DB *mongo.Client
}

func (mongodb *MongoDB) Init() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Minute)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_DB_URI")))
	if err != nil {
		log.Fatalf("Error: %v", err.Error())
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Printf("Error: %v", err.Error())
		log.Fatalf("Error Connecting to Databse...")
	} else {
		log.Printf("Database Connected Successfully")
		mongodb.DB = client
	}
}
func (mongodb *MongoDB) InitializeCollection(collection string) *mongo.Collection {
	return mongodb.DB.Database("ContactKeeper").Collection(collection)
}