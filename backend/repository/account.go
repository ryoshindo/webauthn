package repository

import (
	"context"

	"github.com/ryoshindo/webauthn/backend/api/graph/model"
)

type AccountRepository interface {
	FindByID(context.Context, string) (*model.Account, error)
	FindByEmail(context.Context, string) (*model.Account, error)
	Create(ctx context.Context, account *model.Account) error
}
