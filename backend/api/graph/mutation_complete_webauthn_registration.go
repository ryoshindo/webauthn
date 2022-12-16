package graph

import (
	"context"
	"fmt"

	"github.com/ryoshindo/webauthn/backend/api/graph/model"
)

func (r *mutationResolver) CompleteWebauthnRegistration(ctx context.Context, input model.CompleteWebauthnRegistrationInput) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}
