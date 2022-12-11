package graph

import "github.com/ryoshindo/webauthn/backend/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	accountRepo repository.AccountRepository
}

func NewResolver(accountRepo repository.AccountRepository) *Resolver {
	return &Resolver{
		accountRepo: accountRepo,
	}
}
