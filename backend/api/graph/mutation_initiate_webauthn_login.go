package graph

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ryoshindo/webauthn/backend/api/graph/model"
	"github.com/ryoshindo/webauthn/backend/api/graph/session"
	"github.com/ryoshindo/webauthn/backend/repository"
)

func (r *mutationResolver) InitiateWebauthnLogin(ctx context.Context, input model.InitiateWebauthnLoginInput) (string, error) {
	repoOpts := repository.Options{
		WithWebauthnCredential: true,
	}
	account, err := r.accountRepo.FindByEmail(ctx, input.Email, repoOpts)
	fmt.Println("acccccount", account.WebauthnCredentials)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ACCOUNT_NOT_FOUND")
	}

	options, data, err := r.webAuthn.BeginLogin(account)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("FAILED_INITIATE_WEBAUTHN_LOGIN")
	}

	account.WebauthnLogin.Challenge = data.Challenge
	if err := session.UpdateSession(ctx, account); err != nil {
		fmt.Println(err.Error())
		return "", errors.New("FAILED_UPDATE_ACCOUNT_SESSION")
	}

	s, err := json.Marshal(options)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("FAILED_MARSHAL_WEBAUTHN_OPTIONS")
	}

	return string(s), nil
}
