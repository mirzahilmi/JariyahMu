package model

import "github.com/gofiber/fiber/v2"

type Paseto struct {
	Encrypter func(id string) (string, error)
	Decrypter func(c *fiber.Ctx) error
}
