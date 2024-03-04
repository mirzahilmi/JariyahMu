package rest

import (
	"github.com/MirzaHilmi/JariyahMu/internal/app/usecase"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/model"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	usc usecase.UserUsecaseItf
}

func NewUserHandler(usc usecase.UserUsecaseItf, router fiber.Router) {
	userHandler := UserHandler{usc}
	router = router.Group("/auth")

	router.Post("/signup", userHandler.SignUp)
}

func (h *UserHandler) SignUp(c *fiber.Ctx) error {
	var payloadUser model.CreateUserRequest
	if err := c.BodyParser(&payloadUser); err != nil {
		return err
	}

	token, err := h.usc.RegisterUser(c.Context(), payloadUser)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
