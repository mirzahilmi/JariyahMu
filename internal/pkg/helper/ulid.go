package helper

import (
	"math/rand"
	"time"

	"github.com/MirzaHilmi/JariyahMu/internal/pkg/pool"
	"github.com/oklog/ulid/v2"
)

func NewULID() (ulid.ULID, error) {
	entropy := pool.EntropyPool.Get().(*rand.Rand)
	ms := ulid.Timestamp(time.Now())

	id, err := ulid.New(ms, entropy)
	if err != nil {
		return ulid.ULID{}, err
	}

	return id, nil
}
