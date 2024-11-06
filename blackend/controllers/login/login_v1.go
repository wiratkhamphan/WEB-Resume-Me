package login

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/wiratkhamphan/WEBResumeMe/config/database"
	"github.com/wiratkhamphan/WEBResumeMe/models"
	"golang.org/x/crypto/bcrypt"
)

// const jwtSecret = "infinitas"

func Login_v1(c *fiber.Ctx) error {
	request := models.LoginRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}
	if request.Username == "" || request.Password == "" {
		return fiber.ErrUnprocessableEntity
	}

	db, err := database.Connect()
	if err != nil {
		log.Println("Database connection error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database connection error"})
	}
	defer db.Close()

	// Fetch user details from the database
	user := models.User{}
	query := "SELECT id, username, password FROM user_login WHERE username = ?"
	err = db.Get(&user, query, request.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Incorrect username or password  ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง",
			})
		}
		log.Println("Query error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Query error",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Incorrect username or password")
	}
	// Generate JWT token
	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Println("JWT generation error:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Could not generate token")
	}

	return c.JSON(fiber.Map{
		"jwtToken": token,
		"status":   "ok",
		"massage":  "Login successful",
	})
}
