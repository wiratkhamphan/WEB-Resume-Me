package controllers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/wiratkhamphan/WEB-Resume-Me.git/database"
	"github.com/wiratkhamphan/WEB-Resume-Me.git/models"
	"golang.org/x/crypto/bcrypt"
)

// JWT Secret key (load from environment variable for security)
var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // Ensure this environment variable is set

// Custom claims struct for JWT
type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour).Truncate(jwt.TimePrecision) // Truncate for compatibility

	claims := jwt.RegisteredClaims{
		Issuer:    "",
		Subject:   "",
		Audience:  []string{},
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		NotBefore: jwt.NewNumericDate(time.Now()), // Set NotBefore to current time
		IssuedAt:  jwt.NewNumericDate(time.Now()), // Set IssuedAt to current time
		ID:        "",
	}

	// Sign the token with your secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return signedToken,
		nil
}

// เข้าสู่ระบบ
func Login(c *fiber.Ctx) error {
	var data map[string]string

	// ตรวจสอบการดึงข้อมูลจาก body
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse body",
			"error":   err.Error(),
		})
	}

	var user models.User

	// ตรวจสอบว่าผู้ใช้มีอยู่ในฐานข้อมูลหรือไม่
	if err := database.DB.Where("username = ?", data["username"]).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// ตรวจสอบรหัสผ่าน
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	// สร้าง JWT token
	token, err := GenerateJWT(user.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create token",
		})
	}

	// ถ้าการเข้าสู่ระบบสำเร็จ สามารถตอบกลับข้อความได้ที่นี่
	return c.JSON(fiber.Map{
		"status":      "ok",
		"message":     "Login successful",
		"accessToken": token, // ส่งกลับ access token
		"username":    user.Username,
	})
}

// ลงทะเบียน
func Register(c *fiber.Ctx) error {
	var data map[string]string

	// ตรวจสอบการดึงข้อมูลจาก body
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse body",
			"error":   err.Error(),
		})
	}

	// เข้ารหัสผ่าน
	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
			"error":   err.Error(),
		})
	}

	user := models.User{
		Username: data["username"],
		Password: string(password),
	}

	// บันทึกผู้ใช้ในฐานข้อมูล
	result := database.DB.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create user",
			"error":   result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Registration successful",
	})
}
