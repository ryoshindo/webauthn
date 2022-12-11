package graph

import (
	"context"

	"github.com/ryoshindo/webauthn/backend/api/graph/session"
)

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	session.Remove(ctx)
	return true, nil
}
