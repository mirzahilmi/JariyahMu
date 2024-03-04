package helper

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
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
