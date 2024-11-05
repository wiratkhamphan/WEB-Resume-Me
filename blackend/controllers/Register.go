package controllers

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wiratkhamphan/WEBResumeMe/config/database"
	"github.com/wiratkhamphan/WEBResumeMe/models"
	"golang.org/x/crypto/bcrypt"
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

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Signup(c *fiber.Ctx) error {
	// Parse the signup request JSON
	request := SignupRequest{}
	if err := c.BodyParser(&request); err != nil {
		// Log and return error if JSON parsing fails
		log.Println("Error parsing JSON:", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON format")
	}

	// Check if required fields are empty
	if request.Username == "" || request.Password == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Username and password are required")
	}

	// Connect to the database
	db, err := database.Connect()
	if err != nil {
		log.Println("Database connection error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database connection error"})
	}
	defer db.Close()

	// Check if the username already exists
	var existingUser models.User
	err = db.QueryRow("SELECT id, username FROM user_login WHERE username = ?", request.Username).Scan(&existingUser.Id, &existingUser.Username)
	if err == nil {
		// If no error occurs, it means the username already exists
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error":    "Username already exists",
			"username": request.Username,
		})
	} else if err != sql.ErrNoRows {
		// If an unexpected error occurs
		log.Println("Database query error:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Database query error")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Password encryption error")
	}

	// Insert the user record into the database
	query := "INSERT INTO user_login (username, password) VALUES (?, ?)"
	result, err := db.Exec(query, request.Username, string(hashedPassword))
	if err != nil {
		log.Println("Database insert error:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create user")
	}

	// Retrieve the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error retrieving user ID")
	}

	// Return the created user (without the password for security)
	user := models.User{
		Id:       int(id),
		Username: request.Username,
		Password: string(hashedPassword),
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
