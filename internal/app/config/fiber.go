package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewFiber(viper *viper.Viper) *fiber.App {
	fiber := fiber.New(fiber.Config{
		AppName:          viper.GetString("APP_NAME"),
		DisableKeepalive: viper.GetBool("ALLOW_KEEP_ALIVE"),
		Prefork:          viper.GetBool("SOCKET_SHARDING"),
		StrictRouting:    viper.GetBool("STRICT_ROUTING"),
		ErrorHandler:     NewErrorHandler(),
		RequestMethods:   []string{fiber.MethodHead, fiber.MethodGet, fiber.MethodPost, fiber.MethodPut, fiber.MethodPatch, fiber.MethodDelete},
	})

	return fiber
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
