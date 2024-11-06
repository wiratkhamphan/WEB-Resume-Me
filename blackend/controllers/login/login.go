package login

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/wiratkhamphan/WEBResumeMe/config/database"
	"github.com/wiratkhamphan/WEBResumeMe/models"
	"golang.org/x/crypto/bcrypt"
)

// JWT Secret key (load from environment variable for security)
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Ensure this environment variable is set

// Custom claims struct for JWT
type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateJWT creates a new JWT token for the given username
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token will expire in 24 hours

	claims := CustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "YourAppName",    // Set your app name here
			Subject:   username,         // Set subject as the username
			Audience:  []string{"user"}, // Define audience
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Sign the token with your secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Login handles user authentication and token generation
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
	defer db.Close() // Ensure to close the database connection

	var storedPassword string
	err = db.QueryRow("SELECT password FROM user_login WHERE username = ?", userLogin.Username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid credentials",
			})
		}
		log.Println("Query error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Query error",
		})
	}

	// ถ้ารหัสpassword ไม่เท่ากับที่ Query มาให้แจ้งเตือน "message": "Invalid credentials"
	// if userLogin.Password != storedPassword {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"message": "Invalid credentials",
	// 	})
	// }
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(userLogin.Password))
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Incorrect username or password")
	}
	// Generate JWT token
	token, err := GenerateJWT(userLogin.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create token",
		})
	}

	return c.JSON(fiber.Map{
		"status":      "ok",
		"message":     "Login successful",
		"accessToken": token,
	})
}
