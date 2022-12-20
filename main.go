package main

import (
	"github.com/anderws/go-products/src/database"
	"github.com/anderws/go-products/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()
	database.AutoMigrate()	

	app := fiber.New()
	routes.Setup(app)

	app.Listen(":8000")
}