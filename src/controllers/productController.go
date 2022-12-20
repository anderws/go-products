package controllers

import (
	"fmt"
	"strconv"

	"github.com/anderws/go-products/src/database"
	"github.com/anderws/go-products/src/models"
	"github.com/gofiber/fiber/v2"
)

func Insert(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Erro ao fazer parser body: ", err)
		return err
	}

	priceStr := data["price"]
	price, _ := strconv.ParseInt(priceStr, 10, 64)
	
	product := models.Product {
		Name: data["name"],
		Price: price,
		Category: data["category"],
	}

	database.DB.Create(&product)
	return c.JSON(product.Id)
}

func GetById(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")  
	if err != nil {
		return err
	}
	var product models.Product
	database.DB.Where("id = ?", id).First(&product)

	if product.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Product not found",
		})
	}
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"product": product.Name,
	})
}

func Delete(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")  
	if err != nil {
		return err
	}

    var product models.Product
	database.DB.Where("id = ?", id).First(&product)

	if product.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	database.DB.Delete(&models.Product{}, id)
	c.Status(fiber.StatusNoContent)
	return c.JSON(fiber.Map{
		"message": "Product deleted",
	})
}