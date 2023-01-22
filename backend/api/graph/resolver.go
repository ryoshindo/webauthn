package graph

import (
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/ryoshindo/webauthn/backend/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	accountRepo            repository.AccountRepository
	webauthnCredentialRepo repository.WebauthnCredentialRepository
	webAuthn               *webauthn.WebAuthn
}

func NewResolver(
	accountRepo repository.AccountRepository,
	webauthnCredentialRepo repository.WebauthnCredentialRepository,
) *Resolver {
	webAuthn, _ := webauthn.New(&webauthn.Config{
		RPDisplayName: "Ryo Shindo",
		RPID:          "localhost",
		RPOrigin:      "http://localhost:3333", // FIXME
	})

	return &Resolver{
		accountRepo:            accountRepo,
		webauthnCredentialRepo: webauthnCredentialRepo,
		webAuthn:               webAuthn,
	}
}
