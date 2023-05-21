package router

import (
	"github.com/gofiber/fiber/v2"
	instructorRoutes "github.com/vincemoke66/keyper-api/internals/routes/instructor"
	studentRoutes "github.com/vincemoke66/keyper-api/internals/routes/student"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api")

	studentRoutes.SetupStudentRoutes(api)
	instructorRoutes.SetupStudentRoutes(api)
}
