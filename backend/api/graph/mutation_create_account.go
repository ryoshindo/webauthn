package graph

import (
	"context"
	"errors"

	"github.com/ryoshindo/webauthn/backend/api/graph/model"
	"github.com/ryoshindo/webauthn/backend/api/graph/session"
	m "github.com/ryoshindo/webauthn/backend/model"
)

func (r *mutationResolver) CreateAccount(ctx context.Context, input model.CreateAccountInput) (bool, error) {
	account := m.NewAccount()
	account.Email = input.Email
	account.UserName = input.UserName
	if err := r.accountRepo.Create(ctx, account); err != nil {
		return false, errors.New("INAPPROPRIATE_ACCOUNT_INPUT")
	}

	session.CreateSession(ctx, account)

	return true, nil
}
