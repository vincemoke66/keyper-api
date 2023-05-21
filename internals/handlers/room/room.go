package roomHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vincemoke66/keyper-api/database"
	"github.com/vincemoke66/keyper-api/internals/model"
)

// GetRooms func gets all existing rooms
// @Description Get all existing rooms
// @Tags Room
// @Accept json
// @Produce json
// @Success 200 {array} model.Room
// @router /api/room [get]
func GetRooms(c *fiber.Ctx) error {
	db := database.DB
	var rooms []model.Room

	// find all rooms in the database
	db.Find(&rooms)

	// If no room is present return an error
	if len(rooms) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Rooms data found", "data": nil})
	}

	// Else return rooms
	return c.JSON(fiber.Map{"status": "success", "message": "Rooms Found", "data": rooms})
}

// CreateRoom func create a room
// @Description Create a Room
// @Tags Room
// @Accept json
// @Produce json
// @Param name body string true "name"
// @Param floor body int true "floor"
// @Param building body Building true "building"
// @Success 200 {object} model.Room
// @router /api/room [post]
func CreateRoom(c *fiber.Ctx) error {
	db := database.DB
	room := new(model.Room)

	// Parse the body to the room object
	err := c.BodyParser(room)
	// Return parse error if any
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Return invalid school_id if empty or null
	if room.Name == uuid.Nil.String() || room.Name == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Name", "data": err})
	}

	// Create a temporary room data
	var storedRoom model.Room

	// Find the room with the given name
	db.Find(&storedRoom, "name = ?", room.Name)

	// If room exists with given name, return an error
	if storedRoom.ID != uuid.Nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Room with the same name already exist.", "data": nil})
	}

	// Add a uuid to the new room
	room.ID = uuid.New()

	// Create the Room
	err = db.Create(&room).Error
	// return error if encountered
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create room", "data": err})
	}

	// Return the created room
	return c.JSON(fiber.Map{"status": "success", "message": "Room created", "data": room})
}

// GetRoom func get one room by name
// @Description Get one room by name
// @Tags Room
// @Accept json
// @Produce json
// @Success 200 {object} model.Room
// @router /api/room/{name} [get]
func GetRoom(c *fiber.Ctx) error {
	db := database.DB
	var room model.Room

	// Read the param room_name
	room_name := c.Params("name")

	// Find the room with the given room name
	db.Find(&room, "name = ?", room_name)

	// If no such room present, return an error
	if room.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Room not found", "data": nil})
	}

	// Return the room with the specified name
	return c.JSON(fiber.Map{"status": "success", "message": "Room Found", "data": room})
}

// UpdateRoom update a room by name
// @Description Update a Room by name
// @Tags Room
// @Accept json
// @Produce json
// @Param name body string true "name"
// @Param floor body int true "floor"
// @Param building body Building true "building"
// @Success 200 {object} model.Room
// @router /api/room/{name} [put]
func UpdateRoom(c *fiber.Ctx) error {
	// Create a struct for updating only writable values
	type updateRoom struct {
		Name     string         `json:"name"`
		Floor    int            `json:"floor"`
		Building model.Building `json:"building"`
	}

	db := database.DB
	var room model.Room

	// Read the param room_name
	room_name := c.Params("name")

	// Find the room with the given room name param
	db.Find(&room, "name = ?", room_name)

	// If no such room, return an error
	if room.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Room not found", "data": nil})
	}

	// Store the body containing the updated data
	var updateRoomData updateRoom
	err := c.BodyParser(&updateRoomData)

	// Return parsing error if encountered
	if err != nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Edit the room
	room.Name = updateRoomData.Name
	room.Floor = updateRoomData.Floor
	room.Building = updateRoomData.Building

	// Save the Changes
	db.Save(&room)

	// Return the updated room
	return c.JSON(fiber.Map{"status": "success", "message": "Room Updated", "data": room})
}

// DeleteRoom delete a room by name
// @Description Delete a Room by name
// @Tags Room
// @Accept json
// @Produce json
// @Success 200
// @router /api/room/{name} [delete]
func DeleteRoom(c *fiber.Ctx) error {
	db := database.DB
	var room model.Room

	// Read the param room_name
	room_name := c.Params("name")

	// Find the room with the given room_name param
	db.Find(&room, "name = ?", room_name)

	// If no such room present return an error
	if room.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Room not found", "data": nil})
	}

	// Delete the room
	err := db.Delete(&room, "name = ?", room_name).Error

	// Return error if encountered
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to delete room", "data": nil})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Room Deleted"})
}
