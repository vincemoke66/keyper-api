package router

import (
	"github.com/gofiber/fiber/v2"
	attendanceRoutes "github.com/vincemoke66/keyper-api/internals/routes/attendance"
	buildingRoutes "github.com/vincemoke66/keyper-api/internals/routes/building"
	instructorRoutes "github.com/vincemoke66/keyper-api/internals/routes/instructor"
	keyRoutes "github.com/vincemoke66/keyper-api/internals/routes/key"
	recordRoutes "github.com/vincemoke66/keyper-api/internals/routes/record"
	roomRoutes "github.com/vincemoke66/keyper-api/internals/routes/room"
	scheduleRoutes "github.com/vincemoke66/keyper-api/internals/routes/schedule"
	studentRoutes "github.com/vincemoke66/keyper-api/internals/routes/student"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api")

	studentRoutes.SetupStudentRoutes(api)
	instructorRoutes.SetupStudentRoutes(api)
	buildingRoutes.SetupStudentRoutes(api)
	roomRoutes.SetupStudentRoutes(api)
	keyRoutes.SetupStudentRoutes(api)
	recordRoutes.SetupStudentRoutes(api)
	attendanceRoutes.SetupStudentRoutes(api)
	scheduleRoutes.SetupStudentRoutes(api)
}
