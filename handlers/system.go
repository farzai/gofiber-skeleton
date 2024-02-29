package handlers

import "github.com/gofiber/fiber/v2"

func ApplySystemRoutes(app *fiber.App) {
	systemHandler := NewSystemHandler()

	app.Get("/api/v1/health-check", systemHandler.HealthCheck)
}

func NewSystemHandler() *SystemHandler {
	return &SystemHandler{}
}

type SystemHandler struct {}

func (h *SystemHandler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
	})
}
