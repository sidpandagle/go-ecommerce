package main

import (
	"github.com/gofiber/fiber/v2"

	"ecommerce/database"
	"ecommerce/routes"
)

func main() {
	app := fiber.New()
	database.Connect()
	routes.SetupRoutes(app)
	app.Listen(":8080")
}
