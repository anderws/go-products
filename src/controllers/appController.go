package controllers

import (
	"github.com/anderws/go-products/src/database"
	"github.com/anderws/go-products/src/models"
	"github.com/gofiber/fiber/v2"
)

func AppHealth(c *fiber.Ctx) {
	var count int64
	database.DB.Table("products").Count(&count)
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"product": product.Name,
	})
}

func AppUp(c *fiber.Ctx) {
	return c.Status(fiber.StatusOK).Send("Up")
}
