package router

import (
	"api/handler"

	"github.com/gofiber/fiber/v2"
)

// Setup our router
func SetupOrdersRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Order Group
	groupOrder := api.Group("/order")

	// Get ticket informations
	groupOrder.Get("/", handler.GetAllOrders)
	groupOrder.Get("/get", handler.GetSingleOrder)

	// Create Order
	createOrder := groupOrder.Group("/create")
	createOrder.Post("/", handler.CreateOrder)

	// Create Order
	confirmOrder := groupOrder.Group("/done")
	confirmOrder.Post("/", handler.ConfirmPayment)

	// Cancel Order
	cancelOrder := groupOrder.Group("/cancel")
	cancelOrder.Post("/", handler.CancelPayment)

	// Delete Order
	// DO NOT USE THIS ACTION FOR THE SAKE OF HUMANITY
	// ONLY FOR TESTING PURPOSE
	deleteOrder := groupOrder.Group("/delete")
	deleteOrder.Post("/", handler.DeleteOrder)

}
