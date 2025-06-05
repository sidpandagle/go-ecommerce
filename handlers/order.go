package handlers

import (
	"github.com/gofiber/fiber/v2"

	"ecommerce/database"
	"ecommerce/dto"
	"ecommerce/models"
)

func CreateOrder(c *fiber.Ctx) error {
	order := new(models.Order)
	if err := c.BodyParser(order); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	database.DB.Create(&order)
	return c.JSON(order)
}

func GetAllOrder(c *fiber.Ctx) error {
	var orders []models.Order
	database.DB.Preload("User").Find(&orders)
	return c.JSON(orders)
}

func GetOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	var order models.Order
	result := database.DB.Preload("User").First(&order, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "order not found"})
	}
	return c.JSON(order)
}

func UpdateOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	var order models.Order
	if err := database.DB.First(&order, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "order not found"})
	}
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	database.DB.Save(&order)
	return c.JSON(order)
}

func DeleteOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Order{}, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "order not found"})
	}
	return c.SendStatus(204)
}

func PlaceOrder(c *fiber.Ctx) error {
	orderRequest := new(dto.Order)
	if err := c.BodyParser(orderRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	order := new(models.Order)
	order.UserID = orderRequest.UserID

	var total float64
	for _, orderItem := range orderRequest.OrderItems {
		var product models.Product
		if err := database.DB.First(&product, orderItem.ProductID).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
		}

		itemTotal := product.Price * float64(orderItem.Quantity)
		total += itemTotal

		order.OrderItems = append(order.OrderItems, models.OrderItem{
			OrderID:   orderItem.OrderID,
			UserID:    orderItem.UserID,
			ProductID: orderItem.ProductID,
			Quantity:  orderItem.Quantity,
		})
	}
	order.Total = total
	database.DB.Create(&order)
	return c.JSON(order)
}
