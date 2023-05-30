package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vincemoke66/keyper-api/database"
	"github.com/vincemoke66/keyper-api/router"
)

func main() {
	app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	// Setup the router
	router.SetupRoutes(app)

	app.Listen("192.168.147.91:8080")
}
