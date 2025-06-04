package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	UserID    uint
	User      User
	ProductID uint
	Product   Product
	Quantity  int
}
