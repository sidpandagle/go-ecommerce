package handlers

import (
	"github.com/gofiber/fiber/v2"

	"ecommerce/database"
	"ecommerce/models"
)

func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	database.DB.Create(&product)
	return c.JSON(product)
}

func GetAllProduct(c *fiber.Ctx) error {
	var products []models.Product
	database.DB.Preload("Category").Find(&products)
	return c.JSON(products)
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	result := database.DB.Preload("Category").First(&product, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "product not found"})
	}
	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "product not found"})
	}
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	database.DB.Save(&product)
	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Product{}, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "product not found"})
	}
	return c.SendStatus(204)
}
