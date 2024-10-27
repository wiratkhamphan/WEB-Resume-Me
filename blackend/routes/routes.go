package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiratkhamphan/WEBResumeMe/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Ensure that login route is POST
	app.Post("/login", controllers.Login)
}
