package scheduleRoutes

import (
	"github.com/gofiber/fiber/v2"
	scheduleHandler "github.com/vincemoke66/keyper-api/internals/handlers/schedule"
)

func SetupStudentRoutes(router fiber.Router) {
	schedule := router.Group("/schedule")

	// Create a schedule
	schedule.Post("/", scheduleHandler.CreateSchedule)
	// Read all rooms
	schedule.Get("/", scheduleHandler.GetSchedules)
}
