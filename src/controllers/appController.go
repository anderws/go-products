package controllers

import (
	"github.com/anderws/go-products/src/database"
	"github.com/gofiber/fiber/v2"
)

func AppHealth(c *fiber.Ctx) error {
	var count int64
	database.DB.Table("products").Select("count(id)").Count(&count)
	c.Status(fiber.StatusOK)
	return c.SendString("Ok")
}

func AppUp(c *fiber.Ctx) error {
	c.Status(fiber.StatusOK)
	return c.SendString("Up")
}
