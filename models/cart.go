package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID    uint
	User      User
	ProductID uint
	Product   Product
	Quantity  int
}
