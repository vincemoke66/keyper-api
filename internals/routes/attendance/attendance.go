package attendanceRoutes

import (
	"github.com/gofiber/fiber/v2"
	attendanceHandler "github.com/vincemoke66/keyper-api/internals/handlers/attendance"
)

func SetupStudentRoutes(router fiber.Router) {
	attendance := router.Group("/attendance")

	// Create a attendance
	attendance.Post("/", attendanceHandler.CreateAttendance)
	// Read all rooms
	attendance.Get("/", attendanceHandler.GetAttendance)
}
