package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Description string
	Price       float64
	Stock       int
	CategoryID  uint
	Category    Category
}
