package handlers

import (
	"github.com/gofiber/fiber/v2"

	"ecommerce/database"
	"ecommerce/models"
)

func CreateOrderItem(c *fiber.Ctx) error {
	orderItems := new(models.OrderItem)
	if err := c.BodyParser(orderItems); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	database.DB.Create(&orderItems)
	return c.JSON(orderItems)
}

func GetAllOrderItem(c *fiber.Ctx) error {
	var orderItemss []models.OrderItem
	database.DB.Preload("User").Preload("Product").Preload("Product.Category").Find(&orderItemss)
	return c.JSON(orderItemss)
}

func GetOrderItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var orderItems models.OrderItem
	result := database.DB.Preload("User").Preload("Product").Preload("Product.Category").First(&orderItems, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "orderItems not found"})
	}
	return c.JSON(orderItems)
}

func UpdateOrderItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var orderItems models.OrderItem
	if err := database.DB.First(&orderItems, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "orderItems not found"})
	}
	if err := c.BodyParser(&orderItems); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	database.DB.Save(&orderItems)
	return c.JSON(orderItems)
}

func DeleteOrderItem(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.OrderItem{}, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "orderItems not found"})
	}
	return c.SendStatus(204)
}
