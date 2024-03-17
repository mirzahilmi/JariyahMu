package test

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/MirzaHilmi/JariyahMu/internal/pkg/helper"
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

func mustULID() string {
	id, err := helper.ULID()
	if err != nil {
		panic(err)
	}

	return id
}
