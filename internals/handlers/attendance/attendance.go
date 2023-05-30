package attendanceHandler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vincemoke66/keyper-api/database"
	"github.com/vincemoke66/keyper-api/internals/model"
	"gorm.io/gorm"
)

// GetAttendance func gets all existing attendance
// @Description Get all existing attendance
// @Tags Attendance
// @Accept json
// @Produce json
// @Success 200 {array} model.Attendance
// @router /api/attendance [get]
func GetAttendance(c *fiber.Ctx) error {
	db := database.DB
	var attendances []model.Attendance

	// find all attendances in the database
	db.Find(&attendances)

	// If no attendance is present return an error
	if len(attendances) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Attendances data found", "data": nil})
	}

	// Else return attendances
	return c.JSON(fiber.Map{"status": "success", "message": "Attendances Found", "data": attendances})
}

// CreateAttendance func creates an attendance
// @Description Creates a Attendance
// @Tags Attendance
// @Accept json
// @Produce json
// @Param room_name body string true "room_name"
// @Param rfid body string true "rfid"
// @Success 200 {object} model.Attendance
// @router /api/attendance [post]
func CreateAttendance(c *fiber.Ctx) error {
	db := database.DB
	attendance := new(model.Attendance)

	type AttendanceToAdd struct {
		RoomName string `json:"room"`
		RFID     string `json:"rfid"`
	}

	attendance_to_add := new(AttendanceToAdd)

	// Parse the json request body to the attendance object
	err := c.BodyParser(attendance_to_add)
	// Return parse error if any
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Create a temporary student data
	var storedStudent model.Student
	db.Find(&storedStudent, "rfid = ?", attendance_to_add.RFID)
	// If student does not exist, return an error
	if storedStudent.ID == uuid.Nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Student does not exist.", "data": nil})
	}

	// Create a temporary room data
	var storedRoom model.Room
	db.Find(&storedRoom, "name = ?", attendance_to_add.RoomName)
	// If room name does not exists, return an error
	if storedRoom.ID == uuid.Nil {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Room does not exist.", "data": nil})
	}

	// Check if the entered room and current time has a schedule
	currentTime := time.Now().Format("15:04:05")
	hasSchedule, scheduleFound, err := CheckSchedule(currentTime, storedRoom.Name)

	if !hasSchedule {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid input.", "data": nil})
	}

	// Validate if not a repeating attendance
	// find all attendances in the database
	var latestAttendance model.Attendance
	currentDate := time.Now().Format("2006-01-02")
	query := db.Where("schedule_id = ? AND DATE(created_at) = ? AND student_id = ?", scheduleFound.ID, currentDate, storedStudent.ID).First(&latestAttendance)
	if query.Error != gorm.ErrRecordNotFound {
		return c.Status(300).JSON(fiber.Map{"status": "error", "message": "Student Already Attended", "data": query.Error})
	}

	// Add a uuid to the new attendance
	attendance.ID = uuid.New()

	attendance.StudentName = storedStudent.LastName + ", " + storedStudent.FirstName
	attendance.Section = storedStudent.Section
	attendance.Course = storedStudent.Course
	attendance.RoomName = attendance_to_add.RoomName
	attendance.Subject = scheduleFound.Subject
	attendance.ScheduleID = scheduleFound.ID
	attendance.StudentID = storedStudent.ID

	// Create the Schedule and return error if encountered
	err = db.Create(&attendance).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create attendance", "data": err})
	}

	// Return the created attendance
	return c.JSON(fiber.Map{"status": "success", "message": "Attendance recorded", "data": attendance})
}

// Function to check if the input matches a schedule
func CheckSchedule(inputTime string, roomName string) (bool, model.Schedule, error) {
	db := database.DB
	var schedule model.Schedule

	// Query to check if there is a matching schedule
	err := db.Where("start_time <= ? AND end_time >= ? AND room_name = ?", inputTime, inputTime, roomName).First(&schedule).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// No matching schedule found
			return false, model.Schedule{}, nil
		}
		// Error occurred during the query
		return false, model.Schedule{}, err
	}

	// Matching schedule found
	return true, schedule, nil
}
