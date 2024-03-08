package test

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/MirzaHilmi/JariyahMu/internal/pkg/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSignUp(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	err := truncate("Users")
	require.Nil(err)

	payloadStruct := model.CreateUserRequest{
		FullName:             "John Doe",
		Email:                "john.doe@gmail.com",
		Password:             "12345678",
		PasswordConfirmation: "12345678",
	}
	payloadJSON, err := json.Marshal(&payloadStruct)
	require.Nil(err)
	payload := strings.NewReader(string(payloadJSON))

	req := httptest.NewRequest(fiber.MethodPost, "/api/v1/auth/signup", payload)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	resRaw, err := app.Test(req)
	require.Nil(err)

	defer resRaw.Body.Close()

	assert.Equal(fiber.StatusCreated, resRaw.StatusCode)

	resBody, err := io.ReadAll(resRaw.Body)
	require.Nil(err)

	var res map[string]any
	assert.Nil(json.Unmarshal(resBody, &res))

	token, err := pasetoo.Decode(res["token"].(string))
	require.Nil(err)

	iss, err := token.GetIssuer()
	require.Nil(err)

	issued, err := token.GetIssuedAt()
	require.Nil(err)

	expire, err := token.GetExpiration()
	require.Nil(err)

	assert.Equal(viper.GetString("APP_HOST"), iss)
	assert.Less(issued, time.Now())
	assert.Greater(expire, time.Now())
}
