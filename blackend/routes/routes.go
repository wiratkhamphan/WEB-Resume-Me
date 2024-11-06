package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiratkhamphan/WEBResumeMe/controllers"
	login "github.com/wiratkhamphan/WEBResumeMe/controllers/login"
)

// Set up routes

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//  URL : port/login
	app.Post("/login", login.Login)
	app.Post("/login_v1", login.Login_v1)
	app.Get("/api/user", controllers.GetUser)
	app.Post("/api/Register", controllers.Register)
	app.Post("/api/Signup", controllers.Signup)

}
