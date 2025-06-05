package routes

import (
	"ecommerce/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Auth
	api.Post("/register", handlers.Register)
	api.Post("/login", handlers.Login)

	// User
	api.Post("/users", handlers.CreateUser)
	api.Get("/users", handlers.GetAllUser)
	api.Get("/users/:id", handlers.GetUser)
	api.Put("/users/:id", handlers.UpdateUser)
	api.Delete("/users/:id", handlers.DeleteUser)

	// Category
	api.Post("/categories", handlers.CreateCategory)
	api.Get("/categories", handlers.GetAllCategory)
	api.Get("/categories/:id", handlers.GetCategory)
	api.Put("/categories/:id", handlers.UpdateCategory)
	api.Delete("/categories/:id", handlers.DeleteCategory)

	// Product
	api.Post("/products", handlers.CreateProduct)
	api.Get("/products", handlers.GetAllProduct)
	api.Get("/products/:id", handlers.GetProduct)
	api.Put("/products/:id", handlers.UpdateProduct)
	api.Delete("/products/:id", handlers.DeleteProduct)

	// Cart
	api.Post("/carts", handlers.CreateCart)
	api.Get("/carts", handlers.GetAllCart)
	api.Get("/carts/:id", handlers.GetCart)
	api.Put("/carts/:id", handlers.UpdateCart)
	api.Delete("/carts/:id", handlers.DeleteCart)

	// Order
	api.Post("/orders", handlers.CreateOrder)
	api.Get("/orders", handlers.GetAllOrder)
	api.Get("/orders/:id", handlers.GetOrder)
	api.Put("/orders/:id", handlers.UpdateOrder)
	api.Delete("/orders/:id", handlers.DeleteOrder)
	api.Post("/orders/place-order", handlers.PlaceOrder)

	// OrderItem
	api.Post("/order_items", handlers.CreateOrderItem)
	api.Get("/order_items", handlers.GetAllOrderItem)
	api.Get("/order_items/:id", handlers.GetOrderItem)
	api.Put("/order_items/:id", handlers.UpdateOrderItem)
	api.Delete("/order_items/:id", handlers.DeleteOrderItem)
}
