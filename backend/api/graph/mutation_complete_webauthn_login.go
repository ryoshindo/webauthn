package graph

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/ryoshindo/webauthn/backend/api/graph/model"
	"github.com/ryoshindo/webauthn/backend/api/graph/session"
	"github.com/ryoshindo/webauthn/backend/repository"
)

func (r *mutationResolver) CompleteWebauthnLogin(ctx context.Context, input model.CompleteWebauthnLoginInput) (bool, error) {
	repoOpts := repository.Options{
		WithWebauthnCredential: true,
	}
	account, err := r.accountRepo.FindByEmail(ctx, input.Email, repoOpts)
	if err != nil {
		return false, errors.New("ACCOUNT_NOT_FOUND")
	}

	parsedResponse, err := protocol.ParseCredentialRequestResponseBody(strings.NewReader(input.Credential))
	if err != nil {
		fmt.Println(err.Error())
		return false, errors.New("INAPPROPRIATE_WEBAUTHN_LOGIN_CREDENTIAL_INPUT")
	}

	sess := session.GetSession(ctx)
	data := webauthn.SessionData{
		Challenge:        sess.Account.WebauthnLogin.Challenge,
		UserID:           []byte(account.ID),
		UserVerification: protocol.VerificationRequired,
	}
	_, err = r.webAuthn.ValidateLogin(account, data, parsedResponse)
	if err != nil {
		fmt.Println(err.Error())
		return false, errors.New("WEBAUTHN_LOGIN_FAILED")
	}

	return true, nil
}
