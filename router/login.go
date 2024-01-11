package router

import (
	"api/config"
	"api/handler"
	"api/middleware"

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

	// Admin Routes
	admin.Get("/", handler.GetAllAdmin)
	admin.Get("/:id", handler.GetSingleAdmin)
	admin.Post("/", handler.CreateAdmin)
	admin.Put("/:id", handler.UpdateAdmin)
	admin.Delete("/:id", handler.DeleteAdminByID)
	admin.Post("/auth", handler.AdminLogin)

	// User Group
	promotor := api.Group("/promotor")

	// Admin Routes
	promotor.Get("/", handler.GetAllPromotor)
	promotor.Get("/:id", handler.GetSinglePromotor)
	promotor.Post("/", handler.CreatePromotor)
	promotor.Put("/:id", handler.UpdatePromotor)
	promotor.Delete("/:id", handler.DeletePromotorByID)
	promotor.Post("/auth", handler.PromotorLogin)

	// User Group
	login := api.Group("/auth")

	// Request token
	login.Post("/", handler.Login)

	// Get UserUID
	getid := api.Group("/auth/id")
	jwt := middleware.NewAuthMiddleware(config.Config("AUTH_SECRET"))
	getid.Get("/", jwt, handler.GetUUID)
}
