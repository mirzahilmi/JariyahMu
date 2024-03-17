package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/MirzaHilmi/JariyahMu/internal/app/usecase"
	"github.com/MirzaHilmi/JariyahMu/internal/pkg/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSignUp(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	err := truncate("Users")
	require.Nil(err)

	cases := []struct {
		user           model.CreateUserRequest
		title          string
		expectedStatus int
		expectedResp   []byte
	}{
		{
			user: model.CreateUserRequest{
				FullName:             "John Doe",
				Email:                "john.doe@gmail.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			title:          "Normal",
			expectedStatus: fiber.StatusCreated,
			expectedResp:   []byte(utils.StatusMessage(fiber.StatusCreated)),
		},
		{
			user: model.CreateUserRequest{
				FullName:             "Jean Doe",
				Email:                "john.doe@gmail.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			title:          "EmailAlreadyExist",
			expectedStatus: fiber.StatusConflict,
			expectedResp: mustJSONMarshal(fiber.Map{
				"errors": fiber.Map{"message": usecase.ErrEmailExist.Error()},
			}),
		},
		{
			user: model.CreateUserRequest{
				FullName:             "Jean Doe",
				Email:                "jean.doe@gmail.com",
				Password:             "1234567",
				PasswordConfirmation: "1234567",
			},
			title:          "ValidationError",
			expectedStatus: fiber.StatusBadRequest,
			expectedResp: mustJSONMarshal(fiber.Map{
				"errors": fiber.Map{
					"message":  utils.StatusMessage(fiber.StatusBadRequest),
					"password": "must be atleast 8 characters length",
				},
			}),
		},
	}

	for _, payload := range cases {
		t.Run(payload.title, func(t *testing.T) {
			raw, err := json.Marshal(&payload.user)
			require.Nil(err)
			buff := bytes.NewBuffer(raw)

			req := httptest.NewRequest(fiber.MethodPost, "/api/v1/auth/signup", buff)
			req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
			res, err := app.Test(req, -1)
			require.Nil(err)
			defer res.Body.Close()
			body, err := io.ReadAll(res.Body)
			require.Nil(err)

			assert.Equal(payload.expectedStatus, res.StatusCode)
			assert.Equal(payload.expectedResp, body)
		})
	}
}

func TestUserVerification(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	user := model.CreateUserRequest{
		FullName:             "John Doe",
		Email:                "john.doe@gmail.com",
		Password:             "12345678",
		PasswordConfirmation: "12345678",
	}

	raw, err := json.Marshal(user)
	require.Nil(err)
	buff := bytes.NewBuffer(raw)

	req := httptest.NewRequest(fiber.MethodPost, "/api/v1/auth/signup", buff)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	_, err = app.Test(req, -1)
	require.Nil(err)

	var attempt model.UserVerificationResource
	err = db.Get(&attempt, `
		SELECT ID, UserID, Token FROM UserVerifications 
		WHERE UserID = (
			SELECT ID FROM Users WHERE Email = ? LIMIT 1
		)
		LIMIT 1;`, user.Email)
	require.Nil(err)

	cases := []struct {
		title          string
		attemptID      string
		token          string
		expectedStatus int
		expectedResp   []byte
	}{
		{
			title:          "WrongToken",
			attemptID:      attempt.ID,
			token:          "123456",
			expectedStatus: fiber.StatusBadRequest,
			expectedResp: mustJSONMarshal(fiber.Map{
				"errors": fiber.Map{"message": usecase.ErrVerificationNotExist.Error()},
			}),
		},
		{
			title:          "WrongID",
			attemptID:      "AAAAAAAAAAAAAAAAAAAAAAAAAA",
			token:          attempt.Token,
			expectedStatus: fiber.StatusBadRequest,
			expectedResp: mustJSONMarshal(fiber.Map{
				"errors": fiber.Map{"message": usecase.ErrVerificationNotExist.Error()},
			}),
		},
		{
			title:          "Normal",
			attemptID:      attempt.ID,
			token:          attempt.Token,
			expectedStatus: fiber.StatusOK,
			expectedResp:   []byte(utils.StatusMessage(fiber.StatusOK)),
		},
	}

	for _, payload := range cases {
		t.Run(payload.title, func(t *testing.T) {
			req = httptest.NewRequest(fiber.MethodGet, fmt.Sprintf("/api/v1/auth/verify/%s?t=%s", payload.attemptID, payload.token), nil)
			res, err := app.Test(req, -1)
			require.Nil(err)
			defer res.Body.Close()
			body, err := io.ReadAll(res.Body)
			require.Nil(err)

			assert.Equal(payload.expectedStatus, res.StatusCode)
			assert.Equal(payload.expectedResp, body)
		})
	}

}
