package helper

import (
	"math/rand"
	"time"

	"github.com/MirzaHilmi/JariyahMu/internal/pkg/pool"
	"github.com/oklog/ulid/v2"
)

func NewULID() (ulid.ULID, error) {
	ms := ulid.Timestamp(time.Now())
	entropy := pool.EntropyPool.Get().(*rand.Rand)

	id, err := ulid.New(ms, entropy)
	if err != nil {
		return ulid.ULID{}, err
	}
	pool.EntropyPool.Put(entropy)

	return id, nil
}
