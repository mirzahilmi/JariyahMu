package rest

import (
	"github.com/MirzaHilmi/JariyahMu/internal/app/usecase"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	usecase   usecase.UserUsecaseItf
	validator *validator.Validate
}

func RegisterUserHandler(
	usecase usecase.UserUsecaseItf,
	validator *validator.Validate,
	router fiber.Router,
) {
	userHandler := UserHandler{usecase, validator}
	router = router.Group("/auth")

	router.Post("/signup", userHandler.signUp)
	router.Post("/login", userHandler.login)
	router.Get("/verify/:id", userHandler.verify)
}

func (h *UserHandler) signUp(c *fiber.Ctx) error {
	var payloadUser model.CreateUserRequest
	if err := c.BodyParser(&payloadUser); err != nil {
		return err
	}

	if err := h.validator.Struct(&payloadUser); err != nil {
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

	return c.SendStatus(fiber.StatusOK)
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": signedToken})
}
