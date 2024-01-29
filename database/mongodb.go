package database /*

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"tr05-server/log"
)

// Connect to a MongoDB database
func Connect(uri string, database string) (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	// Check if the database exists
	db := client.Database(database)
	err = db.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	log.Info("Connected to MongoDB")

	return client, nil
}

// Close the MongoDB database connection
func Close(client *mongo.Client) {
	err := client.Disconnect(context.Background())
	if err != nil {
		log.Error("Error closing MongoDB connection:", err)
	}
}
*/
