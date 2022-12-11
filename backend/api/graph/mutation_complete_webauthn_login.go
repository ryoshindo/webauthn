package graph

import (
	"context"
	"fmt"

	"github.com/ryoshindo/webauthn/backend/api/graph/model"
)

func (r *mutationResolver) CompleteWebauthnLogin(ctx context.Context, input model.CompleteWebauthnLoginInput) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}
