package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wiratkhamphan/WEBResumeMe/config/database"
	"github.com/wiratkhamphan/WEBResumeMe/routes"
)

func main() {
	fmt.Println("dev code app running...")

	// Connect to the database
	database.Connect()

	// Create a new Fiber app
	app := fiber.New()

	// Use CORS middleware
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Set("Access-Control-Allow-Credentials", "true")

		if c.Method() == "OPTIONS" {
			c.SendStatus(fiber.StatusNoContent)
			return nil
		}

		return c.Next()
	})

	// Set up routes
	routes.Setup(app)

	// รันที่ port 8080
	app.Listen(":8080")

	// Start the server
	// go func() {
	// 	if err := app.Listen(":8080"); err != nil {
	// 		log.Fatalf("Error Starimg Server: %v", err)
	// 	}
	// }()

	// // Handle graceful shutdown
	// stop := make(chan os.Signal, 1)
	// signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// <-stop // Wait for interrupt signal
	// fmt.Println("Shutting down server...")
	// if err := app.Shutdown(); err != nil {
	// 	log.Fatalf("Error shutting down server: %v", err)
	// }
	// fmt.Println("Server stopped gracefully.")
}
