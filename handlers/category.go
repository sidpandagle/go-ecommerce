package handlers

import (
	"github.com/gofiber/fiber/v2"

	"ecommerce/database"
	"ecommerce/models"
)

func CreateCategory(c *fiber.Ctx) error {
	category := new(models.Category)
	if err := c.BodyParser(category); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	database.DB.Create(&category)
	return c.JSON(category)
}

func GetAllCategory(c *fiber.Ctx) error {
	var categories []models.Category
	database.DB.Find(&categories)
	return c.JSON(categories)
}

func GetCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	result := database.DB.First(&category, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "category not found"})
	}
	return c.JSON(category)
}

func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "category not found"})
	}
	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	database.DB.Save(&category)
	return c.JSON(category)
}

func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Category{}, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "category not found"})
	}
	return c.SendStatus(204)
}
