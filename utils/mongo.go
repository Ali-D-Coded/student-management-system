package utils

import (
	"context"
	"student-management-system/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateUniqueIndexes creates unique indexes for the specified fields in the given collection
func CreateUniqueIndexes(collectionName string, fields ...string) error {
	collection := database.GetCollection(collectionName)
	var indexes []mongo.IndexModel

	for _, field := range fields {
		index := mongo.IndexModel{
			Keys:    bson.D{{Key: field, Value: 1}},
			Options: options.Index().SetUnique(true),
		}
		indexes = append(indexes, index)
	}

	_, err := collection.Indexes().CreateMany(context.Background(), indexes)
	return err
}