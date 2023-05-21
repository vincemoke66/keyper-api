package router

import (
	"github.com/gofiber/fiber/v2"
	buildingRoutes "github.com/vincemoke66/keyper-api/internals/routes/building"
	instructorRoutes "github.com/vincemoke66/keyper-api/internals/routes/instructor"
	roomRoutes "github.com/vincemoke66/keyper-api/internals/routes/room"
	studentRoutes "github.com/vincemoke66/keyper-api/internals/routes/student"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api")

	studentRoutes.SetupStudentRoutes(api)
	instructorRoutes.SetupStudentRoutes(api)
	buildingRoutes.SetupStudentRoutes(api)
	roomRoutes.SetupStudentRoutes(api)
}
