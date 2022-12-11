package graph

import (
	"context"
	"errors"

	"github.com/ryoshindo/webauthn/backend/api/graph/model"
	"github.com/ryoshindo/webauthn/backend/api/graph/session"
)

func (r *mutationResolver) CreateAccount(ctx context.Context, input model.CreateAccountInput) (*model.Account, error) {
	account := model.Account{
		Email: input.Email,
		UserName: input.UserName,
	}
	if err := r.accountRepo.Create(ctx, &account); err != nil {
		return nil, errors.New("INAPPROPRIATE_ACCOUNT_INPUT")
	}

	session.Set(ctx, &account)

	return &account, nil
}
