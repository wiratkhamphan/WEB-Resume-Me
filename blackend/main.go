package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiratkhamphan/WEB-Resume-Me.git/controllers"
	"github.com/wiratkhamphan/WEB-Resume-Me.git/database"
)

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		c.Set("Access-Control-Allow-Credentials", "true")

		if c.Method() == "OPTIONS" {
			c.SendStatus(fiber.StatusNoContent)
			return nil
		}

		return c.Next()
	})
	// Connect to Database
	database.ConnectDB()

	// Routes
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)

	app.Listen(":3000") // รันที่ port 3000
}
