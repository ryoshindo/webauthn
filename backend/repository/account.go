package repository

import (
	"context"

	"github.com/ryoshindo/webauthn/backend/model"
)

type Options struct {
	WithWebauthnCredential bool
}

type AccountRepository interface {
	FindByID(context.Context, string) (*model.Account, error)
	FindByEmail(context.Context, string, Options) (*model.Account, error)
	Create(context.Context, *model.Account) error
	CreateWebauthnCredential(context.Context, *model.Account, *model.WebauthnCredential) error
}
