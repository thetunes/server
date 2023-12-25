package handler

import (
	"api/database"
	"api/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Create a Admin
func CreateAdmin(c *fiber.Ctx) error {
	db := database.DB.Db
	admin := new(model.Admin)
	// Store the body in the Admin and return error if encountered
	err := c.BodyParser(admin)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	err = db.Create(&admin).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create Admin", "data": err})
	}
	// Return the created Admin
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Admin has created", "data": admin})
}

// Get All Admin from db
func GetAllAdmin(c *fiber.Ctx) error {
	db := database.DB.Db
	var admin []model.Admin
	// find all Admin in the database
	db.Find(&admin)
	// If no Admin found, return an error
	if len(admin) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Admin not found", "data": nil})
	}
	// return Admin
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Admin Found", "data": admin})
}

// GetSingleAdmin from db
func GetSingleAdmin(c *fiber.Ctx) error {
	db := database.DB.Db
	// get id params
	id := c.Params("id")
	var admin model.Admin
	// find single admin in the database by id
	db.Find(&admin, "id = ?", id)
	if admin.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Admin not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Admin Found", "data": admin})
}

// update a user in db
func UpdateAdmin(c *fiber.Ctx) error {
	type updateAdmin struct {
		Username string `json:"username"`
	}
	db := database.DB.Db
	var admin model.Admin
	// get id params
	id := c.Params("id")
	// find single admin in the database by id
	db.Find(&admin, "id = ?", id)
	if admin.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Admin not found", "data": nil})
	}
	var updateAdminData updateAdmin
	err := c.BodyParser(&updateAdminData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	admin.Username = updateAdminData.Username
	// Save the Changes
	db.Save(&admin)
	// Return the updated admin
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Admin Found", "data": admin})
}

// delete Admin in db by ID
func DeleteAdminByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var admin model.Admin
	// get id params
	id := c.Params("id")
	// find single Admin in the database by id
	db.Find(&admin, "id = ?", id)
	if admin.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Admin not found", "data": nil})
	}
	err := db.Delete(&admin, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete Admin", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Admin deleted"})
}
