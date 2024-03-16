package config

import (
	"errors"

	"github.com/MirzaHilmi/JariyahMu/internal/pkg/helper"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
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
		var apiErr *response.Error
		if errors.As(err, &apiErr) {
			return ctx.Status(apiErr.Code).JSON(fiber.Map{
				"errors": fiber.Map{"message": apiErr.Error()},
			})
		}

		if validationErr, ok := err.(validator.ValidationErrors); ok {
			fieldErr := fiber.Map{}
			for _, e := range validationErr {
				fieldErr[e.Field()] = helper.ValidationErrMsg(e)
			}
			fieldErr["message"] = utils.StatusMessage(fiber.StatusBadRequest)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errors": fieldErr,
			})
		}

		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": fiber.Map{"message": utils.StatusMessage(fiberErr.Code)},
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": fiber.Map{"message": utils.StatusMessage(fiber.StatusInternalServerError)},
		})
	}
}
