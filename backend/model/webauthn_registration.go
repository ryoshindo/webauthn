package model

import "time"

type WebauthnRegistration struct {
	ID        string    `bun:"id,pk"`
	CreatedAt time.Time `bun:"created_at"`
	UpdatedAt time.Time `bun:"updated_at"`
	Challenge string    `bun:"challenge"`

	Account    *Account            `bun:"rel:belongs-to,join:account_id=id"`
	Credential *WebauthnCredential `bun:"rel:has-one,join:registration_id=id"`
}

type WebauthnRegistrationList []WebauthnRegistration

func NewWebauthnRegistration() *WebauthnRegistration {
	return &WebauthnRegistration{
		ID: NewULIDString(),
	}
}
