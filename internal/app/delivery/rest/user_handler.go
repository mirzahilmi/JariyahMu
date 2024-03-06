package rest

import (
	"github.com/MirzaHilmi/JariyahMu/internal/app/usecase"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/model"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	usecase usecase.UserUsecaseItf
}

func RegisterUserHandler(usecase usecase.UserUsecaseItf, router fiber.Router) {
	userHandler := UserHandler{usecase}
	router = router.Group("/auth")

	router.Post("/signup", userHandler.signUp)
}

func (h *UserHandler) signUp(c *fiber.Ctx) error {
	var payloadUser model.CreateUserRequest
	if err := c.BodyParser(&payloadUser); err != nil {
		return err
	}

	token, err := h.usecase.RegisterUser(c.Context(), payloadUser)
	if err != nil {
		return err
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"token": token,
	})
}
