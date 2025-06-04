package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"ecommerce/models"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	MigrateModels()
}

func MigrateModels() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.Cart{},
		&models.Order{},
	)
	if err != nil {
		log.Fatal("Failed to migrate DB:", err)
	}
}
