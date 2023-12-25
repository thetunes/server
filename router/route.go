package router

import (
	"api/handler"

	"github.com/gofiber/fiber/v2"
)

// Setup our router
func SetupUserRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User Group
	user := api.Group("/user")

	// User Routes
	user.Get("/", handler.GetAllUsers)
	user.Get("/:id", handler.GetSingleUser)
	user.Post("/", handler.CreateUser)
	user.Put("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUserByID)

	// User Group
	admin := api.Group("/admin")

	// User Routes
	admin.Get("/", handler.GetAllAdmin)
	admin.Get("/:id", handler.GetSingleAdmin)
	admin.Post("/", handler.CreateAdmin)
	admin.Put("/:id", handler.UpdateAdmin)
	admin.Delete("/:id", handler.DeleteAdminByID)

}
