package main

import (
	"api/config"
	"api/database"
	"api/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	router.SetupUserRoutes(app)
	router.SetupTicketRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	port := config.Config("PORT")

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
