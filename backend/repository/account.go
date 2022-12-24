package repository

import (
	"context"

	"github.com/ryoshindo/webauthn/backend/model"
)

type AccountRepository interface {
	FindByID(context.Context, string) (*model.Account, error)
	FindByEmail(context.Context, string) (*model.Account, error)
	Create(context.Context, *model.Account) error
}
