package repository

import (
	"context"

	"github.com/ryoshindo/webauthn/backend/model"
)

type WebauthnCredentialRepository interface {
	FindByID(context.Context, string) (*model.WebauthnCredential, error)
	Create(context.Context, *model.WebauthnCredential) error
}
