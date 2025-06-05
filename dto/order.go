package dto

type Order struct {
	UserID     uint
	Total      float64
	OrderItems []OrderItem
}
