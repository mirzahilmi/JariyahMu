package test

import (
	"fmt"
	"sync"
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
