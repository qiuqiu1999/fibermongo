package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qiuqiu1999/fibermongo/controller"
)

func UserRoute(app *fiber.App) {
	// All routes related to users comes here
	//app.Get("/", func(c *fiber.Ctx) error {
	//	return c.JSON(&fiber.Map{"welcome": "Hello from Fiber + MongoDB"})
	//})
	app.Post("/user", controller.CreateUser)

	app.Get("/user/:userId", controller.GetAUser)
	app.Put("/user/:userId", controller.EditAUser)
	app.Delete("/user/:userId", controller.DeleteAUser)
	app.Get("/users", controller.GetAllUsers)
}
