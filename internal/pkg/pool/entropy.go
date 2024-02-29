package pool

import (
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

var EntropyPool = sync.Pool{
	New: func() any {
		entropy := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
		return &entropy
	},
}
