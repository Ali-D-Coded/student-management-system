package users

import (
	"fmt"
	"student-management-system/database"
	"student-management-system/models"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var validate = validator.New()


func GetAllUsers(c *fiber.Ctx) error {



	coll := database.GetCollection("users")

	// return all users
	filter := bson.M{}
	opts := options.Find().SetSkip(0).SetLimit(100)

	// find all users
	cursor, err := coll.Find(c.Context(), filter, opts)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	// convert cursor to slice
	users := make([]models.Users, 0)
	if err = cursor.All(c.Context(), &users); err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	// return users
	return c.Status(200).JSON(users)
}



func CreateUser(c *fiber.Ctx) error {
	// Initialize user DTO
	var userDTO CreateUserDTO

	// Parse request body into user DTO
	if err := c.BodyParser(&userDTO); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Validate the user DTO
	if err := validate.Struct(&userDTO); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Validation failed", "details": err.Error()})
	}

	hashed, err := database.HashPassword(userDTO.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}

	// Convert the DTO to the Users model
	user := models.Users{
		ID:           primitive.NewObjectID(),
		Username:     userDTO.Username,
		PasswordHash: hashed,
		Email:        userDTO.Email,
		Role:         userDTO.Role,
		RefreshToken: userDTO.RefreshToken,
		CreatedAt:    primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:    primitive.NewDateTimeFromTime(time.Now()),
	}

	// Insert user into the database
	coll := database.GetCollection("users")
	_, err = coll.InsertOne(c.Context(), user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user", "details": err.Error()})
	}

	return c.Status(201).JSON(user)
}


func UpdateUser(c *fiber.Ctx) error {

//userInterface := c.Locals("user")

//user, ok := userInterface.(jwt.MapClaims)
//if !ok {
//fmt.Println("Not fopubnd")
//}
//userID := user["user_id"].(string) // Assuming user_id is a string
//fmt.Println("User ID:", userID)

	// Get user ID from the request parameters
	userID := c.Params("id")

	// Convert userID to MongoDB's ObjectID type
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	
	// Initialize user DTO
	var updateUser UpdateUserDTO

	// Parse request body into user DTO
	if err := c.BodyParser(&updateUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	
	// Validate the user DTO
	if err := validate.Struct(&updateUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Validation failed", "details": err.Error()})
	}
	
	hash, err := database.HashPassword(updateUser.Password)
	
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid password"})
	}

	// Create the update document
	update := bson.M{
		"$set": bson.M{
			"username":     updateUser.Username,
			"email":        updateUser.Email,
			"password": hash,
			"role":         updateUser.Role,
			"updatedAt":    primitive.NewDateTimeFromTime(time.Now()),
		},
	}

	// Get the collection
	coll := database.GetCollection("users")

	// Update the user in the database
	result, err := coll.UpdateOne(c.Context(), bson.M{"_id": objID}, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update user", "details": err.Error()})
	}

	// Check if the user was found and updated
	if result.MatchedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "User updated successfully"})
}


func UpdatePassword(c *fiber.Ctx) error {
	userID := c.Params("id")
// Convert userID to MongoDB's ObjectID type
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	fmt.Println(objID)


	return c.Status(200).JSON(fiber.Map{"message": "Password updated successfully"})
}