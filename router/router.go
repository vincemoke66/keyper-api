package router

import (
	"github.com/gofiber/fiber/v2"
	studentRoutes "github.com/vincemoke66/keyper-api/internals/routes/student"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api")

	studentRoutes.SetupStudentRoutes(api)
}
