package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiratkhamphan/WEB-Resume-Me.git/database"
	"github.com/wiratkhamphan/WEB-Resume-Me.git/models"
	"golang.org/x/crypto/bcrypt"
)

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

	// ถ้าการเข้าสู่ระบบสำเร็จ สามารถตอบกลับข้อความได้ที่นี่
	return c.JSON(fiber.Map{
		"message": "Login successful",
		// คุณสามารถส่งข้อมูลผู้ใช้หรือ token ที่ใช้ในการเข้าถึง API ได้
		"username": user.Username,
	})
}
