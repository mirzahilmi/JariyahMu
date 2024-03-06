package helper

import (
	"time"

	"github.com/oklog/ulid/v2"
	"golang.org/x/exp/rand"
)

func ULID() (string, error) {
	ms := ulid.Timestamp(time.Now())
	entropy := entropyPool.Get().(*rand.Rand)

	id, err := ulid.New(ms, entropy)
	entropyPool.Put(entropy)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
