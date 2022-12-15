package graph

import (
	"context"
	"errors"

	"github.com/ryoshindo/webauthn/backend/api/graph/model"
	"github.com/ryoshindo/webauthn/backend/api/graph/session"
)

func (r *queryResolver) Viewer(ctx context.Context) (*model.Account, error) {
	account := session.Account(ctx)
	if account == nil {
		return nil, errors.New("UNAUTHORIZED")
	}
	
	return account, nil
}
