package rest

import "github.com/gofiber/fiber/v2"

func RegisterUtilsHandler(app fiber.Router) {
	app.Get("/health", healthCheck)
	app.Get("/status", healthCheck)
}

func healthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "Healthy",
	})
}
