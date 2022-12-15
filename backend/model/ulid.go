package model

import (
	"crypto/rand"
	"sync"
	"time"

	"github.com/oklog/ulid/v2"
)

var (
	defaultEntropySource *ulid.MonotonicEntropy
	entropySourceMtx     sync.Mutex
)

func init() {
	defaultEntropySource = ulid.Monotonic(rand.Reader, 0)
}

func NewULID() ulid.ULID {
	entropySourceMtx.Lock()
	defer entropySourceMtx.Unlock()
	return ulid.MustNew(ulid.Timestamp(time.Now()), defaultEntropySource)
}

func NewULIDString() string {
	return NewULID().String()
}
