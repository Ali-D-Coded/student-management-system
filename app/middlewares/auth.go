package middlewares

import (
	"fmt"
	"strings"
	"student-management-system/utils"

	"github.com/gofiber/fiber/v2"
)

var jwtHandler *utils.JwtHandler

func JwtMiddleware(c *fiber.Ctx) error {

	jwtHandler = utils.NewJwtHandler()
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing Authorization header",
		})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := jwtHandler.VerifyToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	fmt.Println(claims["role"].(string))

	// Add claims to the context for use in subsequent handlers
	c.Locals("user", claims)
	return c.Next()
}