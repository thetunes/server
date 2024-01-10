package handler

import (
	"api/database"
	userCheck "api/helper/username"
	"api/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Create a admin
func CreateAdmin(c *fiber.Ctx) error {
	db := database.DB.Db
	admin := new(model.Admin)

	// Parse the request body into the admin struct
	if err := c.BodyParser(admin); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	// Check if the username is already registered
	if userCheck.IsUsernameTaken(db, admin.Username) {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "username is already registered", "data": nil})
	}

	// Create the admin
	err := db.Create(&admin).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create admin", "data": err})
	}

	// Return the created admin
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Admin has been created", "data": admin})
}

// Get All Admins from db
func GetAllAdmin(c *fiber.Ctx) error {
	db := database.DB.Db
	var admins []model.Admin
	// find all admins in the database
	db.Find(&admins)
	// If no admin found, return an error
	if len(admins) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Admins not found", "data": nil})
	}
	// return admins
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Admins Found", "data": admins})
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

// update a admin in db
func UpdateAdmin(c *fiber.Ctx) error {
	type updateAdmin struct {
		username string `json:"username"`
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
	admin.Username = updateAdminData.username
	// Save the Changes
	db.Save(&admin)
	// Return the updated admin
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "admins Found", "data": admin})
}

// delete admin in db by ID
func DeleteAdminByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var admin model.Admin
	// get id params
	id := c.Params("id")
	// find single admin in the database by id
	db.Find(&admin, "id = ?", id)
	if admin.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Admin not found", "data": nil})
	}
	err := db.Delete(&admin, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete admin", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Admin deleted"})
}
