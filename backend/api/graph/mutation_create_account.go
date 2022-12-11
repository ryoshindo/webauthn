package graph

import (
	"context"
	"fmt"

	"github.com/ryoshindo/webauthn/backend/api/graph/model"
)

func (r *mutationResolver) CreateAccount(ctx context.Context, input model.CreateAccountInput) (*model.Account, error) {
	panic(fmt.Errorf("not implemented"))
}
