package handler

import (
	"api/database"
	"api/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// Create a new order
func CreateOrder(c *fiber.Ctx) error {
	db := database.DB.Db
	order := new(model.TicketsOrder)

	// Extract the user information from the JWT token stored in the context
	user, ok := c.Locals("user").(*jtoken.Token)
	if !ok || user == nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "User information not found in the context"})
	}

	// Now you can safely use the user object
	claims := user.Claims.(jtoken.MapClaims)

	// Use the 'ID' claim as the user identifier
	userID, err := uuid.Parse(claims["ID"].(string))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to parse user ID from token"})
	}

	// Generate a random ID for the order
	order.ID = uuid.New().String()

	// Store the body in the order and return an error if encountered
	err = c.BodyParser(order)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	// Set the UserID field with the retrieved user ID
	order.UserID = userID.String()

	err = db.Create(&order).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create order", "data": err})
	}

	// Return the created order
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Order has created", "data": order})
}

// Get All Tickets from db
func GetAllOrders(c *fiber.Ctx) error {
	db := database.DB.Db
	var orders []model.TicketsOrder
	// find all tickets in the database
	db.Find(&orders)
	// If no Orders found, return an error
	if len(orders) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Orders not found", "data": nil})
	}
	// return Order data
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Orders Found", "data": orders})
}

// GetSingleTicket gets a single ticket from the database by ID
func GetSingleOrder(c *fiber.Ctx) error {
	db := database.DB.Db
	// get id params
	id := c.QueryInt("id", 0)
	var order model.TicketsOrder
	// find single ticket in the database by id
	db.Find(&order, "id = ?", id)
	if order.ID == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Order not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Order Found", "data": order})
}

// delete ticket in db by ID
func DeleteOrder(c *fiber.Ctx) error {
	db := database.DB.Db
	var order model.TicketsOrder
	// get id params
	id := c.QueryInt("id", 0)
	// find single ticket in the database by id
	db.Find(&order, "id = ?", id)
	if order.ID == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Order not found", "data": nil})
	}
	err := db.Delete(&order, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete order", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Rrder deleted"})
}

// Set status true if payment received
func ConfirmPayment(c *fiber.Ctx) error {
	db := database.DB.Db
	// get id params
	id := c.Query("id", "")
	var order model.TicketsOrder
	// find single order in the database by id
	db.Find(&order, "id = ?", id)
	if order.ID == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Order not found", "data": nil})
	}

	// Update the status to true
	order.Status = "true"
	err := db.Save(&order).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to update order status", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Order status updated to true", "data": order})
}

// Set status to false incase something happens
func CancelPayment(c *fiber.Ctx) error {
	db := database.DB.Db
	// get id params
	id := c.Query("id", "")
	var order model.TicketsOrder
	// find single order in the database by id
	db.Find(&order, "id = ?", id)
	if order.ID == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Order not found", "data": nil})
	}

	// Update the status to false
	order.Status = "false"
	err := db.Save(&order).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to update order status", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Order status updated to false", "data": order})
}

// CountAllOrders counts the total number of orders, orders with no status, orders with true status, and orders with false status
func CountAllOrders(c *fiber.Ctx) error {
	db := database.DB.Db
	var totalCount, noStatusCount, trueStatusCount, falseStatusCount int64

	// Count total orders
	err := db.Model(&model.TicketsOrder{}).Count(&totalCount).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to count all orders", "data": err})
	}

	// Count orders with no status
	err = db.Model(&model.TicketsOrder{}).Where("status = ?", "").Count(&noStatusCount).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to count orders with no status", "data": err})
	}

	// Count orders with true status
	err = db.Model(&model.TicketsOrder{}).Where("status = ?", "true").Count(&trueStatusCount).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to count orders with true status", "data": err})
	}

	// Count orders with false status
	err = db.Model(&model.TicketsOrder{}).Where("status = ?", "false").Count(&falseStatusCount).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to count orders with false status", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":        "success",
		"message":       "Count of all orders",
		"total_orders":  totalCount,
		"nostat_orders": noStatusCount,
		"true_orders":   trueStatusCount,
		"false_orders":  falseStatusCount,
	})
}

// CountOrdersForUser counts the total number of orders, orders with no status, orders with true status, and orders with false status for a specific user ID
func CountOrdersForUser(c *fiber.Ctx) error {
	db := database.DB.Db
	userID := c.Params("userid")
	var totalCount, noStatusCount, trueStatusCount, falseStatusCount int64

	// Count total orders for the user
	err := db.Model(&model.TicketsOrder{}).Where("user_id = ?", userID).Count(&totalCount).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to count all orders for the user", "data": err})
	}

	// Count orders with no status for the user
	err = db.Model(&model.TicketsOrder{}).Where("user_id = ? AND status = ?", userID, "").Count(&noStatusCount).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to count orders with no status for the user", "data": err})
	}

	// Count orders with true status for the user
	err = db.Model(&model.TicketsOrder{}).Where("user_id = ? AND status = ?", userID, "true").Count(&trueStatusCount).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to count orders with true status for the user", "data": err})
	}

	// Count orders with false status for the user
	err = db.Model(&model.TicketsOrder{}).Where("user_id = ? AND status = ?", userID, "false").Count(&falseStatusCount).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to count orders with false status for the user", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":        "success",
		"message":       "Count of all orders",
		"total_orders":  totalCount,
		"nostat_orders": noStatusCount,
		"true_orders":   trueStatusCount,
		"false_orders":  falseStatusCount,
	})
}
