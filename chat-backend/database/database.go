package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// DBClient is the MongoDB client
var DBClient *mongo.Client

// Connect establishes a connection to MongoDB
func Connect(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return err
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	DBClient = client
	if err := EnsureUserIndexes(ctx); err != nil {
		_ = client.Disconnect(ctx)
		DBClient = nil
		return err
	}

	log.Println("Connected to MongoDB!")
	return nil
}

// GetCollection returns a handle to a MongoDB collection
func GetCollection(dbName string, collName string) *mongo.Collection {
	return DBClient.Database(dbName).Collection(collName)
}

// Close disconnects from MongoDB
func Close() {
	if DBClient == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := DBClient.Disconnect(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to MongoDB closed.")
}
