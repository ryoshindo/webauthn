package graph

import (
	"context"
	"errors"
	"fmt"

	"github.com/ryoshindo/webauthn/backend/api/graph/model"
	"github.com/ryoshindo/webauthn/backend/api/graph/session"
	m "github.com/ryoshindo/webauthn/backend/model"
)

func (r *mutationResolver) CreateAccount(ctx context.Context, input model.CreateAccountInput) (bool, error) {
	account := m.NewAccount()
	account.Email = input.Email
	account.UserName = input.UserName
	if err := r.accountRepo.Create(ctx, account); err != nil {
		fmt.Println(err)
		return false, errors.New("INAPPROPRIATE_ACCOUNT_INPUT")
	}

	session.Set(ctx, account)

	return true, nil
}
