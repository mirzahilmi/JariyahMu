package helper

import "sync"

type Helper struct {
	EntropyPool sync.Pool
}
