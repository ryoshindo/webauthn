package model

import (
	"time"

	"github.com/ryoshindo/webauthn/backend/util/rand"
)

type Session struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AccountID string    `json:"account_id"`

	ExpiresAt time.Duration
	Token     string
}

const sessionDefaultExpiry = 1 * 24 * time.Hour

func NewSession(accountID string) (*Session, error) {
	maxStringLength := 128
	token, err := rand.GenerateRandomString(maxStringLength)
	if err != nil {
		return nil, err
	}

	return &Session{
		ID:        NewULIDString(),
		AccountID: accountID,
		Token:     token,
		ExpiresAt: sessionDefaultExpiry,
	}, nil
}
