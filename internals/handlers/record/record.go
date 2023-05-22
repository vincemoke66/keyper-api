package recordHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vincemoke66/keyper-api/database"
	"github.com/vincemoke66/keyper-api/internals/model"
)

// GetAllRecords func gets all records
// @Description Gets all records
// @Tags Record
// @Accept json
// @Produce json
// @Success 200 {array} model.Record
// @router /api/record [get]
func GetAllRecords(c *fiber.Ctx) error {
	db := database.DB
	var records []model.Record

	// find all records
	db.Find(&records)

	// If no record is present return an error
	if len(records) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Records data found", "data": nil})
	}

	// Else return records
	return c.JSON(fiber.Map{"status": "success", "message": "Records Found", "data": records})
}

// CreateRecord func creates a record
// @Description Creates a Record
// @Tags Record
// @Accept json
// @Produce json
// @Param type body string true "type"
// @Param student_school_id body string true "student_school_id"
// @Param room_name body string true "room_name"
// @Success 200 {object} model.Record
// @router /api/record [post]
func CreateRecord(c *fiber.Ctx) error {
	db := database.DB
	record := new(model.Record)

	type RecordToAdd struct {
		Type     model.RecordType `json:"type"`
		SchoolID string           `json:"school_id"`
		RFID     string           `json:"rfid"`
		RoomName string           `json:"room_name"`
	}

	record_to_add := new(RecordToAdd)

	// Parse the body to the key object
	err := c.BodyParser(record_to_add)
	// Return parse error if any
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Return invalid school_id if empty or null
	if record_to_add.SchoolID == uuid.Nil.String() || record_to_add.SchoolID == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid School ID", "data": err})
	}
	// Return invalid room_name if empty or null
	if record_to_add.RoomName == uuid.Nil.String() || record_to_add.RoomName == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid Room Name", "data": err})
	}

	// Create a temporary room data
	var storedStudent model.Student
	db.Find(&storedStudent, "school_id = ?", record_to_add.SchoolID)
	// If room name does not exists, return an error
	if storedStudent.ID == uuid.Nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Student does not exist.", "data": nil})
	}

	// Create a temporary room data
	var storedRoom model.Room
	db.Find(&storedRoom, "name = ?", record_to_add.RoomName)
	// If room name does not exists, return an error
	if storedRoom.ID == uuid.Nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Room does not exist.", "data": nil})
	}

	// Create a temporary key data
	var storedKey model.Key
	// Find the key with the given school_id
	db.Find(&storedKey, "rfid = ?", record_to_add.RFID)
	// If key does not exists, return an error
	if storedKey.ID == uuid.Nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Key does not exist.", "data": nil})
	}

	// Add a uuid to the new key
	record.ID = uuid.New()

	record.Type = record_to_add.Type
	record.StudentID = storedStudent.ID
	record.KeyID = storedKey.ID
	record.RoomID = storedRoom.ID

	// Update key status
	if record.Type == "return" {
		storedKey.Status = "available"
	} else if record.Type == "borrow" {
		storedKey.Status = "borrowed"
	}
	db.Save(storedKey)

	// Create the record and return error if encountered
	err = db.Create(&record).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create record", "data": err})
	}

	// Return the created record
	return c.JSON(fiber.Map{"status": "success", "message": "Record created", "data": record})
}

// GetStudent func get one student by school_id
// @Description Get one student by school_id
// @Tags Student
// @Accept json
// @Produce json
// @Success 200 {object} model.Student
// @router /api/student/{school_id} [get]
// func GetStudent(c *fiber.Ctx) error {
// 	db := database.DB
// 	var student model.Student

// 	// Read the param school_id
// 	school_id := c.Params("school_id")

// 	// Find the student with the given school_id
// 	db.Find(&student, "school_id = ?", school_id)

// 	// If no such student present, return an error
// 	if student.ID == uuid.Nil {
// 		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Student not found", "data": nil})
// 	}

// 	// Return the student with the specified school_id
// 	return c.JSON(fiber.Map{"status": "success", "message": "Student Found", "data": student})
// }

// UpdateKey update a key by rfid
// @Description Update a Key by rfid
// @Tags Key
// @Accept json
// @Produce json
// @Param building_name body string true "building_name"
// @Param room_name body string true "room_name"
// @Param status body string true "status"
// @Success 200 {object} model.Key
// @router /api/key/{rfid} [put]
// func UpdateKey(c *fiber.Ctx) error {
// 	// Create a struct for updating only writable values
// 	type KeyToUpdate struct {
// 		BuildingName string          `json:"building_name"`
// 		RoomName     string          `json:"room_name"`
// 		Status       model.KeyStatus `json:"status"`
// 	}

// 	db := database.DB
// 	var key model.Key

// 	// Read the param school_id
// 	rfid := c.Params("rfid")

// 	// Find the key with the given school_id param
// 	db.Find(&key, "rfid = ?", rfid)

// 	// If no such key, return an error
// 	if key.ID == uuid.Nil {
// 		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Key not found", "data": nil})
// 	}

// 	// Store the body containing the updated data
// 	var key_to_update KeyToUpdate
// 	err := c.BodyParser(&key_to_update)
// 	// Return parsing error if encountered
// 	if err != nil {
// 		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
// 	}

// 	// Create a temporary building data
// 	var storedBuidling model.Building
// 	db.Find(&storedBuidling, "name = ?", key_to_update.BuildingName)
// 	// If building does not exists, return an error
// 	if storedBuidling.ID == uuid.Nil {
// 		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Building does not exist.", "data": nil})
// 	}

// 	// Create a temporary room data
// 	var storedRoom model.Room
// 	db.Find(&storedRoom, "name = ?", key_to_update.RoomName)
// 	// If room does not exists, return an error
// 	if storedRoom.ID == uuid.Nil {
// 		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Room does not exist.", "data": nil})
// 	}

// 	// Edit the key
// 	key.BuildingID = storedBuidling.ID
// 	key.RoomID = storedRoom.ID
// 	key.Status = key_to_update.Status

// 	// Save the Changes
// 	db.Save(&key)

// 	// Return the updated key
// 	return c.JSON(fiber.Map{"status": "success", "message": "Key Updated", "data": key})
// }

// DeleteKey delete a key by rfid
// @Description Delete a Key by rfid
// @Tags Key
// @Accept json
// @Produce json
// @Success 200
// @router /api/key/{rfid} [delete]
// func DeleteKey(c *fiber.Ctx) error {
// 	db := database.DB
// 	var key model.Key

// 	// Read the param key_rfid
// 	key_rfid := c.Params("rfid")

// 	// Find the key with the given rfid param
// 	db.Find(&key, "rfid = ?", key_rfid)

// 	// If no such key present return an error
// 	if key.ID == uuid.Nil {
// 		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Key not found", "data": nil})
// 	}

// 	// Delete the key
// 	err := db.Delete(&key, "rfid = ?", key_rfid).Error

// 	// Return error if encountered
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to delete key", "data": nil})
// 	}

// 	// Return success message
// 	return c.JSON(fiber.Map{"status": "success", "message": "Key Deleted"})
// }
