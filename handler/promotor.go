package handler

import (
	"api/database"
	"api/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"errors"

	"gorm.io/gorm"
)

// Create a promotor
func CreatePromotor(c *fiber.Ctx) error {
	db := database.DB.Db
	promotor := new(model.Promotor)

	// Parse the request body into the promotor struct
	if err := c.BodyParser(promotor); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	// Check if the username is already registered
	if isPromotorUsernameTaken(db, promotor.Username) {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "username is already registered", "data": nil})
	}

	// Create the promotor
	err := db.Create(&promotor).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create promotor", "data": err})
	}

	// Return the created promotor
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Promotor has been created", "data": promotor})
}

// Get All Promotors from db
func GetAllPromotor(c *fiber.Ctx) error {
	db := database.DB.Db
	var promotors []model.Promotor
	// find all promotors in the database
	db.Find(&promotors)
	// If no promotor found, return an error
	if len(promotors) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Promotors not found", "data": nil})
	}
	// return promotors
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Promotors Found", "data": promotors})
}

// GetSinglePromotor from db
func GetSinglePromotor(c *fiber.Ctx) error {
	db := database.DB.Db
	// get id params
	id := c.Params("id")
	var promotor model.Promotor
	// find single promotor in the database by id
	db.Find(&promotor, "id = ?", id)
	if promotor.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Promotor not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Promotor Found", "data": promotor})
}

// update a promotor in db
func UpdatePromotor(c *fiber.Ctx) error {
	type updatePromotor struct {
		username string `json:"username"`
	}
	db := database.DB.Db
	var promotor model.Promotor
	// get id params
	id := c.Params("id")
	// find single promotor in the database by id
	db.Find(&promotor, "id = ?", id)
	if promotor.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Promotor not found", "data": nil})
	}
	var updatePromotorData updatePromotor
	err := c.BodyParser(&updatePromotorData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	promotor.Username = updatePromotorData.username
	// Save the Changes
	db.Save(&promotor)
	// Return the updated promotor
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "promotors Found", "data": promotor})
}

// delete promotor in db by ID
func DeletePromotorByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var promotor model.Promotor
	// get id params
	id := c.Params("id")
	// find single promotor in the database by id
	db.Find(&promotor, "id = ?", id)
	if promotor.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Promotor not found", "data": nil})
	}
	err := db.Delete(&promotor, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete promotor", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Promotor deleted"})
}

func isPromotorUsernameTaken(db *gorm.DB, username string) bool {
	var existingPromotor model.Promotor
	if err := db.Where("username = ?", username).First(&existingPromotor).Error; err != nil {
		// Check if the error is due to the record not being found
		return !errors.Is(err, gorm.ErrRecordNotFound)
	}
	// Promotor with the same username already exists
	return true
}