package handler

import (
	"api/database"
	"api/model"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func CreateTicket(c *fiber.Ctx) error {
	db := database.DB.Db
	ticket := new(model.Ticket)

	// Generate a random ID for the ticket
	ticket.ID = uuid.New().String()

	// Store the body in the ticket and return an error if encountered
	err := c.BodyParser(ticket)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	// Validate the ticket data
	validate := validator.New()
	if err := validate.Struct(ticket); err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Unprocessable Entity", "data": err})
	}

	// Initialize Likes with 0
	ticket.Likes = 0

	err = db.Create(&ticket).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create ticket", "data": err})
	}

	// Return the created ticket
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Ticket has created", "data": ticket})
}

// Get All Tickets from db
func GetAllTicket(c *fiber.Ctx) error {
	db := database.DB.Db
	var tickets []model.Ticket
	// find all tickets in the database
	db.Find(&tickets)
	// If no ticket found, return an error
	if len(tickets) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Tickets not found", "data": nil})
	}
	// return tickets
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Tickets Found", "data": tickets})
}

// GetSingleTicket gets a single ticket from the database by ID
func GetSingleTicket(c *fiber.Ctx) error {
	db := database.DB.Db
	// get id params
	id := c.Params("id")
	var ticket model.Ticket
	// find single ticket in the database by id
	db.Find(&ticket, "id = ?", id)
	if ticket.ID == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Ticket not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Ticket Found", "data": ticket})
}

// UpdateTicket updates a ticket in the database by ID
func UpdateTicket(c *fiber.Ctx) error {
	type updateTicket struct {
		Ticketname string `json:"ticketname"`
	}
	db := database.DB.Db
	var ticket model.Ticket
	// get id params
	id := c.Params("id")
	// find single ticket in the database by id
	db.Find(&ticket, "id = ?", id)
	if ticket.ID == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Ticket not found", "data": nil})
	}
	var updateTicketData updateTicket
	err := c.BodyParser(&updateTicketData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// Save the Changes
	db.Save(&ticket)
	// Return the updated ticket
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Ticket updated", "data": ticket})
}

// delete ticket in db by ID
func DeleteTicket(c *fiber.Ctx) error {
	db := database.DB.Db
	var ticket model.Ticket
	// get id params
	id := c.Params("id")
	// find single ticket in the database by id
	db.Find(&ticket, "id = ?", id)
	if ticket.ID == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Ticket not found", "data": nil})
	}
	err := db.Delete(&ticket, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete ticket", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Ticket deleted"})
}

// IncrementLike increments or decrements the like count for a ticket by ID
func IncrementLike(c *fiber.Ctx) error {
	db := database.DB.Db
	// get id params
	id := c.Params("id")
	var ticket model.Ticket
	// find single ticket in the database by id
	db.Find(&ticket, "id = ?", id)
	if ticket.ID == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Ticket not found", "data": nil})
	}

	// Extract the user information from the JWT token stored in the context
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)

	// Use the 'ID' claim as the user identifier
	userID, err := uuid.Parse(claims["ID"].(string))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to parse user ID from token"})
	}

	// Check if the user's UUID is already in the likers array
	var found bool
	for _, likerID := range ticket.Likers {
		if likerID == userID.String() {
			found = true
			break
		}
	}

	// If not found, add the user's UUID to the likers array
	if !found {
		ticket.Likers = append(ticket.Likers, userID.String())

		// Marshal the Likers field to JSON
		likersJSON, err := json.Marshal(ticket.Likers)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to marshal likers to JSON", "data": err})
		}

		// Explicitly set the Likers field in the model before saving
		err = db.Model(&ticket).Update("Likers", likersJSON).Error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to update likers", "data": err})
		}
	}

	// Return the updated ticket
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Like count updated", "data": ticket})
}
