package graph

import (
	"context"
	"fmt"

	"github.com/ryoshindo/webauthn/backend/api/graph/model"
)

func (r *queryResolver) Account(ctx context.Context) (*model.Account, error) {
	panic(fmt.Errorf("not implemented"))
}
