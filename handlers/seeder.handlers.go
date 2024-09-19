package handlers

import (
	"student-management-system/database"

	"github.com/gofiber/fiber/v2"
)

func SeedData(c *fiber.Ctx) error {
	err := database.DatabaseSeeder()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	return c.Status(200).JSON("database seeded successfully")
}