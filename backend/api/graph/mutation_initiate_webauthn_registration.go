package graph

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/ryoshindo/webauthn/backend/api/graph/session"
)

func (r *mutationResolver) InitiateWebauthnRegistration(ctx context.Context) (string, error) {
	account := session.AccountFromSession(ctx)
	if account == nil {
		return "", errors.New("UNAUTHORIZED")
	}

	registerOptions := func(credCreationOpts *protocol.PublicKeyCredentialCreationOptions) {
		credCreationOpts.CredentialExcludeList = account.CredentialExcludeList()
	}

	options, _, err := r.webAuthn.BeginRegistration(account, registerOptions)
	if err != nil {
		return "", errors.New("FAILED_INITIATE_WEBAUTHN_REGISTRATION")
	}

	s, _ := json.Marshal(options)

	return string(s), nil
}
