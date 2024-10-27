package controllers

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wiratkhamphan/WEBResumeMe/config/database"
	"github.com/wiratkhamphan/WEBResumeMe/models"
)

func Login(c *fiber.Ctx) error {
	var userLogin models.User
	if err := c.BodyParser(&userLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse login data"})
	}
	db, err := database.Connect()
	if err != nil {
		log.Println(err) // Log the error instead of using log.Fatal
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database connection error"})
	}
	var storedPassword string
	err = db.QueryRow("SELECT password FROM user_login WHERE username = ?", userLogin.Username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
		}
		log.Println("Query error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Query error"})
	}

	// Replace the password check with a real verification logic (e.g., bcrypt)
	if userLogin.Password != storedPassword { // Temporary direct comparison
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "Login successful",
	})
}
