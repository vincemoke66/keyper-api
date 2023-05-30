package scheduleHandler

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vincemoke66/keyper-api/database"
	"github.com/vincemoke66/keyper-api/internals/model"
)

// GetSchedules func gets all existing schedules
// @Description Get all existing schedules
// @Tags Schedule
// @Accept json
// @Produce json
// @Success 200 {array} model.Schedule
// @router /api/schedule [get]
func GetSchedules(c *fiber.Ctx) error {
	db := database.DB
	var schedules []model.Schedule

	// find all schedules in the database
	db.Find(&schedules)

	// If no schedule is present return an error
	if len(schedules) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Schedules data found", "data": nil})
	}

	// Else return schedules
	return c.JSON(fiber.Map{"status": "success", "message": "Schedules Found", "data": schedules})
}

func CreateSchedule(c *fiber.Ctx) error {
	db := database.DB
	// Parse JSON request body
	type ScheduleToAdd struct {
		RoomName       string `json:"room_name"`
		StartTime      string `json:"start_time" `
		EndTime        string `json:"end_time" `
		DayOfWeek      string `json:"day"`
		Subject        string `json:"subject"`
		InstructorName string `json:"instructor"`
	}
	var reqBody ScheduleToAdd

	err := c.BodyParser(&reqBody)
	if err != nil {
		// Handle parsing error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err,
		})
	}

	// Create a new instance of the Schedule model
	newSchedule := model.Schedule{
		StartTime:      reqBody.StartTime,
		EndTime:        reqBody.EndTime,
		RoomName:       reqBody.RoomName,
		DayOfWeek:      reqBody.DayOfWeek,
		Subject:        reqBody.Subject,
		InstructorName: reqBody.InstructorName,
	}
	newSchedule.ID = uuid.New()

	// Perform validation on the schedule data if needed
	// ...

	// Save the new schedule to the database
	result := db.Create(&newSchedule)
	if result.Error != nil {
		// Handle database error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create schedule",
		})
	}

	// Return success response
	return c.JSON(fiber.Map{
		"message": "Schedule created successfully",
		"data":    newSchedule,
	})
}

func isValidTimeFormat(timeStr string) bool {
	// Define the regular expression pattern for time in the format HH:MM:SS
	pattern := `^([01]\d|2[0-3]):([0-5]\d):([0-5]\d)$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(timeStr)
}
