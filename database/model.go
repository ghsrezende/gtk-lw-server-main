package database /*

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Create a new document in the database
func Create(client *mongo.Client, database string, collection string, document interface{}) error {
	col := client.Database(database).Collection(collection)
	_, err := col.InsertOne(context.Background(), document)
	return err
}

// Read a document from the database
func Read(client *mongo.Client, database string, collection string, id string) (interface{}, error) {
	col := client.Database(database).Collection(collection)
	doc := col.FindOne(context.Background(), bson.M{"_id": id})
	var document interface{}
	err := doc.Decode(&document)
	return document, err
}

// Update a document in the database
func Update(client *mongo.Client, database string, collection string, id string, document interface{}) error {
	col := client.Database(database).Collection(collection)
	_, err := col.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": document})
	return err
}

// Delete a document from the database
func Delete(client *mongo.Client, database string, collection string, id string) error {
	col := client.Database(database).Collection(collection)
	_, err := col.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
*/
