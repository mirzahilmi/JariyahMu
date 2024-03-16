package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

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
		expectedErr    error
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
			expectedErr:    nil,
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
				"errors": fiber.Map{"message": "email already exist"},
			}),
			expectedErr: nil,
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
			expectedErr: nil,
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

// func TestUserVerification(t *testing.T) {
// 	// assert := assert.New(t)
// 	require := require.New(t)

// 	var user model.UserResource
// 	query := `
// 		SELECT u.ID, u.FullName, u.Email, u.HashedPassword, u.ProfilePicture, u.Active
// 		FROM Users AS u
// 		INNER JOIN UserVerifications AS uv
// 		ON u.ID = uv.UserID
// 		WHERE u.Active = FALSE AND uv.Succeed = FALSE
// 		LIMIT 1;
// 	`

// 	err := db.Get(&user, query)
// 	require.Nil(err)

// }
