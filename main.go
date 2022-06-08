package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qiuqiu1999/fibermongo/route"
)

func main() {
	app := fiber.New()
	route.UserRoute(app)
	app.Listen(":2022")
}
