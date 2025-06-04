package handlers

import (
	"github.com/gofiber/fiber/v2"

	"ecommerce/database"
	"ecommerce/models"
)

func CreateCart(c *fiber.Ctx) error {
	cart := new(models.Cart)
	if err := c.BodyParser(cart); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	database.DB.Create(&cart)
	return c.JSON(cart)
}

func GetAllCart(c *fiber.Ctx) error {
	var carts []models.Cart
	database.DB.Find(&carts)
	return c.JSON(carts)
}

func GetCart(c *fiber.Ctx) error {
	id := c.Params("id")
	var cart models.Cart
	result := database.DB.First(&cart, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "cart not found"})
	}
	return c.JSON(cart)
}

func UpdateCart(c *fiber.Ctx) error {
	id := c.Params("id")
	var cart models.Cart
	if err := database.DB.First(&cart, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "cart not found"})
	}
	if err := c.BodyParser(&cart); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	database.DB.Save(&cart)
	return c.JSON(cart)
}

func DeleteCart(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Cart{}, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "cart not found"})
	}
	return c.SendStatus(204)
}
