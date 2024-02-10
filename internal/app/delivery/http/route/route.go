package route

import (
	"golang-clean-architecture/internal/app/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App               *fiber.App
	UserController    *http.UserController
	ContactController *http.ContactController
	AddressController *http.AddressController
	AuthMiddleware    fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Get("/api/status", func(ctx *fiber.Ctx) error {
		return ctx.JSON(struct {
			Status string `json:"status"`
		}{Status: "Healthy"})
	})

	c.App.Post("/api/users", c.UserController.Register)
	c.App.Post("/api/users/_login", c.UserController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	userEnds := c.App.Group("/api/users", c.AuthMiddleware)
	userEnds.Delete("/", c.UserController.Logout)
	userEnds.Patch("/_current", c.UserController.Update)
	userEnds.Get("/_current", c.UserController.Current)

	contactEnds := c.App.Group("/api/contacts", c.AuthMiddleware)
	contactEnds.Get("/", c.ContactController.List)
	contactEnds.Post("/", c.ContactController.Create)
	contactEnds.Put("/:contactId", c.ContactController.Update)
	contactEnds.Get("/:contactId", c.ContactController.Get)
	contactEnds.Delete("/:contactId", c.ContactController.Delete)

	contactEnds.Get("/:contactId/addresses", c.AddressController.List)
	contactEnds.Post("/:contactId/addresses", c.AddressController.Create)
	contactEnds.Put("/:contactId/addresses/:addressId", c.AddressController.Update)
	contactEnds.Get("/:contactId/addresses/:addressId", c.AddressController.Get)
	contactEnds.Delete("/:contactId/addresses/:addressId", c.AddressController.Delete)
}
