package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiratkhamphan/WEBResumeMe/controllers"
)

// Set up routes

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//  URL : port/login
	app.Post("/login", controllers.Login)
	app.Get("/api/user", controllers.GetUser)
	app.Post("/api/Register", controllers.Register)
}
