package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wiratkhamphan/WEBResumeMe/config/database"
	"github.com/wiratkhamphan/WEBResumeMe/models"
	// "golang.org/x/crypto/bcrypt"
)

// Assuming db is your *sql.DB database connection

// HashPassword hashes the password for security
//
//	func HashPassword(password string) (string, error) {
//		bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//		return string(bytes), err
//	}
func Register(c *fiber.Ctx) error {
	var user models.User

	// Parse JSON body into the User struct
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}
	db, err := database.Connect()
	if err != nil {
		log.Println(err) // Log the error instead of using log.Fatal
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database connection error"})
	}
	defer db.Close() // Ensure to close the database connection
	// Hash the user's password
	// hashedPassword, err := HashPassword(user.Password)
	// if err != nil {
	// 	log.Println("Password hashing error:", err)
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": "Server error",
	// 	})
	// }
	// Insert the new user into the database without hashing the password
	// _, err = db.Exec(query, user.Username, hashedPassword)
	// Insert the new user into the database

	// result, err := db.Exec("INSERT INTO user_login (username, password) VALUES (?, ?)", user.Username, hashedPassword)
	_, err = db.Exec("INSERT INTO user_login (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		log.Println("Database insertion error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not register user",
		})
	}

	// Return a success message, including the newly created user ID
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"status":  "ok",
	})
}
