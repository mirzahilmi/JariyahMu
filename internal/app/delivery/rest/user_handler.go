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
	router.Post("/login", userHandler.login)
	router.Post("/verify/:id", userHandler.verify)
}

func (h *UserHandler) signUp(c *fiber.Ctx) error {
	var payloadUser model.CreateUserRequest
	if err := c.BodyParser(&payloadUser); err != nil {
		return err
	}

	if err := h.usecase.RegisterUser(c.Context(), payloadUser); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *UserHandler) verify(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return ErrMissingID
	}

	token := c.Query("t")
	if token == "" {
		return ErrMissingToken
	}

	if err := h.usecase.VerifyUser(c.Context(), id, token); err != nil {
		return err
	}

	return nil
}

func (h *UserHandler) login(c *fiber.Ctx) error {
	var attempt model.UserLoginAttemptRequest
	if err := c.BodyParser(&attempt); err != nil {
		return err
	}

	signedToken, err := h.usecase.LogUserIn(c.Context(), attempt)
	if err != nil {
		return err
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"token": signedToken})
}
