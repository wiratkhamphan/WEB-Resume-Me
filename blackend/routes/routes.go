package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiratkhamphan/WEBResumeMe/config/database"
	"github.com/wiratkhamphan/WEBResumeMe/controllers"
	getproduct "github.com/wiratkhamphan/WEBResumeMe/controllers/GET_Product"
	login "github.com/wiratkhamphan/WEBResumeMe/controllers/login"
)

// Set up routes

func Setup(app *fiber.App) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//  URL : port/login
	app.Post("/login", login.Login)
	app.Post("/login_v1", login.Login_v1)
	app.Get("/api/user", controllers.GetUser)
	app.Post("/api/Register", controllers.Register)
	app.Post("/api/Signup", controllers.Signup)

	// Product
	// app.Get("/product/list")
	// app.Get("/product/form")
	// app.Post("/product")
	// app.Get("/product/{id}")
	// app.Put("/product/{id}")
	// app.Get("/product/remove/{id}")
	// Initialize the ProductService

	app.Get("/product/list", getproduct.Getporductlist)
}
