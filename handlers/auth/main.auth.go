package auth

import (
	"context"
	"fmt"
	"log"
	"student-management-system/database"
	"student-management-system/models"
	"student-management-system/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// Function to check if the password matches the hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}



// Login function handles user login
func Login(c *fiber.Ctx) error {
	// Get the login DTO from the request body
	dto := new(LoginDTO)


	
	// Validate the request body
	if err := c.BodyParser(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input", "details": err.Error()})
	}
	

	// Fetch the user from the MongoDB collection by username
	user := &models.Users{}
	coll := database.GetCollection("users")
	err := coll.FindOne(context.TODO(), bson.M{"username": dto.Username}).Decode(user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
		}
		log.Printf("Error fetching user from the database: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	fmt.Println(user)

	// Validate the password
	if !CheckPasswordHash(dto.PasswordHash, user.PasswordHash) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
	}

	// Generate a new refresh token
	jwtHandler := utils.NewJwtHandler()
	tokenPayload := map[string]interface{}{
		"user_id": user.ID.Hex(),
		"role": user.Role,
		"username":user.Username,
	}
	tokens, err := jwtHandler.GenerateToken(tokenPayload)
	if err != nil {
		log.Printf("Error generating refresh token: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error generating refresh token"})
	}

	// Update the user document with the new refreshToken
	update := bson.M{
		"$set": bson.M{
			"refreshToken": tokens.Refresh,
			"updatedAt":    primitive.NewDateTimeFromTime(time.Now()),
		},
	}
	_, err = coll.UpdateOne(context.TODO(), bson.M{"_id": user.ID}, update)
	if err != nil {
		log.Printf("Error updating user in the database: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error updating user information"})
	}

	// Return the tokens to the client
	return c.JSON(fiber.Map{
		"tokens":  tokens,
	})
}