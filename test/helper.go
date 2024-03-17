package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"sync"

	"github.com/MirzaHilmi/JariyahMu/internal/pkg/model"
	"github.com/gofiber/fiber/v2"
)

func truncate(names ...string) error {
	var wg sync.WaitGroup
	sigChan, errChan := make(chan bool, 1), make(chan error)

	for _, name := range names {
		wg.Add(1)
		go func(name string) {
			_, err := db.Exec(fmt.Sprintf("DELETE FROM %s", name))
			if err != nil {
				errChan <- err
			}
			wg.Done()
		}(name)
	}
	go func() {
		wg.Wait()
		sigChan <- true
	}()

	select {
	case <-sigChan:
		return nil
	case err := <-errChan:
		return err
	}
}

func mustJSONMarshal(v any) []byte {
	raw, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return raw
}

func storeUser(user model.CreateUserRequest) error {
	raw, err := json.Marshal(user)
	if err != nil {
		return err
	}
	buff := bytes.NewBuffer(raw)

	req := httptest.NewRequest(fiber.MethodPost, "/api/v1/auth/signup", buff)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	_, err = app.Test(req, -1)
	if err != nil {
		return err
	}

	return nil
}
