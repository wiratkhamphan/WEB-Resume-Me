package controllers

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/wiratkhamphan/WEBResumeMe/config/database"
	"github.com/wiratkhamphan/WEBResumeMe/models"
)

func GetUser(c *fiber.Ctx) error {
	// Get the JWT from the Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Missing or malformed token",
		})
	}

	// Parse JWT
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid token",
		})
	}

	// Extract username from token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["username"] == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid token claims",
		})
	}

	username := claims["username"].(string)

	// Fetch user details from the database
	user, err := getUserFromDB(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Database error",
		})
	}

	// Return user data as JSON
	return c.JSON(fiber.Map{
		"status": "ok",
		// "user":   user,
		"user": user.Username,
	})
}

// getUserFromDB retrieves user details from the database by username
func getUserFromDB(username string) (*models.User, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user models.User
	err = db.QueryRow("SELECT username FROM user_login WHERE username = ?", username).Scan(&user.Username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
