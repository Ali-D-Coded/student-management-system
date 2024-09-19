package database

import (
	"context"
	"fmt"
	"log"
	"student-management-system/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Function to generate sample seeder data
func generateSeedUsers() ([]interface{}, error) {
	// Define passwords for seeding
	passwords := []string{"admin@123", "password2", "password3", "password4", "password5", "password6", "password7", "password8", "password9", "password10","1234567"}

	// List of users
	users := []models.Users{
		{ID: primitive.NewObjectID(), Username: "admin", Email: "admin@gmail.com", Role: "admin", RefreshToken: "token0"},
		{ID: primitive.NewObjectID(), Username: "johndoe", Email: "johndoe@example.com", Role: "moderator", RefreshToken: "token1"},
		{ID: primitive.NewObjectID(), Username: "janedoe", Email: "janedoe@example.com", Role: "user", RefreshToken: "token2"},
		{ID: primitive.NewObjectID(), Username: "bobsmith", Email: "bobsmith@example.com", Role: "user", RefreshToken: "token3"},
		{ID: primitive.NewObjectID(), Username: "alicejones", Email: "alicejones@example.com", Role: "moderator", RefreshToken: "token4"},
		{ID: primitive.NewObjectID(), Username: "charliebrown", Email: "charliebrown@example.com", Role: "user", RefreshToken: "token5"},
		{ID: primitive.NewObjectID(), Username: "emilywhite", Email: "emilywhite@example.com", Role: "admin", RefreshToken: "token6"},
		{ID: primitive.NewObjectID(), Username: "davidlee", Email: "davidlee@example.com", Role: "user", RefreshToken: "token7"},
		{ID: primitive.NewObjectID(), Username: "sarahmiller", Email: "sarahmiller@example.com", Role: "moderator", RefreshToken: "token8"},
		{ID: primitive.NewObjectID(), Username: "michaelgreen", Email: "michaelgreen@example.com", Role: "user", RefreshToken: "token9"},
		{ID: primitive.NewObjectID(), Username: "rachelblack", Email: "rachelblack@example.com", Role: "admin", RefreshToken: "token10"},
	}

	// Convert to []interface{} for MongoDB insertion
	var documents []interface{}

	for i := range users {
		// Hash each password
		hashedPassword, err := HashPassword(passwords[i])
		if err != nil {
			return nil, err
		}

		// Assign the hashed password to each user
		users[i].PasswordHash = hashedPassword
		users[i].CreatedAt = primitive.NewDateTimeFromTime(time.Now())
		users[i].UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

		// Append the user to the documents slice
		documents = append(documents, users[i])
	}

	return documents, nil
}

// DatabaseSeeder is now exported so it can be called from other packages
func DatabaseSeeder() error {
	// Generate seeder data
	users, err := generateSeedUsers()
	if err != nil {
		return fmt.Errorf("error generating seed users: %w", err)
	}

	// insert the users into the database
	coll := GetCollection("users")
	insertManyResult, err := coll.InsertMany(context.TODO(), users)
	if err != nil {
		return fmt.Errorf("error inserting seed users: %w", err)
	}

	log.Printf("Inserted documents with IDs: %v\n", insertManyResult.InsertedIDs)
	return nil
}

