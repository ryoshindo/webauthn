package model

import "time"

type WebauthnCredential struct {
	ID          string    `bun:"id,pk"`
	CreatedAt   time.Time `bun:"created_at"`
	UpdatedAt   time.Time `bun:"updated_at"`
	AccountID   string    `bun:"account_id"`
	PublicKeyID []byte    `bun:"public_key_id"`
	PublicKey   []byte    `bun:"public_key"`
	ExternalID  string    `bun:"external_id"`
	SignCount   int64     `bun:"sign_count"`

	Account *Account `bun:"rel:belongs-to,join:account_id=id"`
}

type WebauthnCredentialList []WebauthnCredential

func NewWebauthnCredential() *WebauthnCredential {
	return &WebauthnCredential{
		ID: NewULIDString(),
	}
}
