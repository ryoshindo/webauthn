package graph

import (
	"context"

	"github.com/ryoshindo/webauthn/backend/api/graph/session"
)

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	session.DeleteSession(ctx)
	return true, nil
}
