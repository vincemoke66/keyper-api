package keyRoutes

import (
	"github.com/gofiber/fiber/v2"
	keyHandler "github.com/vincemoke66/keyper-api/internals/handlers/key"
)

func SetupStudentRoutes(router fiber.Router) {
	key := router.Group("/key")

	// Create a key
	key.Post("/", keyHandler.CreateKey)

	// Read all key in a specific building
	key.Get("/:building_name", keyHandler.GetKeysUsingBuildingName)

	// Read a key
	// key.Get("/:name", keyHandler.GetRoom)

	// Update key
	key.Put("/:rfid", keyHandler.UpdateKey)

	// Delete a key
	key.Delete("/:rfid", keyHandler.DeleteKey)
}