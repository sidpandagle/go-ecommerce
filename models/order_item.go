package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID   uint
	UserID    uint
	ProductID uint
	Quantity  int
}
