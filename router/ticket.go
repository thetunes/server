package router

import (
	"api/handler"

	"github.com/gofiber/fiber/v2"
)

// Setup our router
func SetupTicketRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User Group
	ticket := api.Group("/ticket")

	// Get All Tickets
	ticket.Get("/", handler.GetAllTicket)

	// Get All by ID
	ticket.Get("/:id", handler.GetSingleUser)

	// Add Ticket
	ticket.Post("/", handler.CreateTicket)

	// Edit Ticket by ID
	ticket.Put("/:id", handler.UpdateTicket)

	// Delete Ticket by ID
	ticket.Delete("/:id", handler.DeleteTicket)
}
