package routes

import (
	"github.com/anderws/go-products/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Get("/", controllers.AppUp)
	app.Get("/health", controllers.AppHealth)

	api := app.Group("api/v1")

	api.Post("products", controllers.Insert)
	api.Get("products/:id", controllers.GetById)
	api.Delete("products/:id", controllers.Delete)
}
