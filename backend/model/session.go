package model

import (
	"time"

	"github.com/ryoshindo/webauthn/backend/util/rand"
)

type Session struct {
	ID      string  `json:"id"`
	Account Account `json:"account"`

	ExpiresAt time.Duration
	Token     string
}

const sessionDefaultExpiry = 1 * 24 * time.Hour

func NewSession(account Account) (*Session, error) {
	maxStringLength := 128
	token, err := rand.GenerateRandomString(maxStringLength)
	if err != nil {
		return nil, err
	}

	return &Session{
		ID:        NewULIDString(),
		Account:   account,
		Token:     token,
		ExpiresAt: sessionDefaultExpiry,
	}, nil
}
