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
	m "github.com/ryoshindo/webauthn/backend/model"
)

func (r *mutationResolver) CompleteWebauthnRegistration(ctx context.Context, input model.CompleteWebauthnRegistrationInput) (bool, error) {
	account := session.AccountFromSession(ctx)
	if account == nil {
		return false, errors.New("UNAUTHORIZED")
	}

	parsedResponse, err := protocol.ParseCredentialCreationResponseBody(strings.NewReader(input.Credential))
	if err != nil {
		fmt.Println(err.Error())
		return false, errors.New("INAPPROPRIATE_WEBAUTHN_REGISTRATION_CREDENTIAL_INPUT")
	}

	sess := session.GetSession(ctx)
	data := webauthn.SessionData{
		Challenge:        sess.Account.WebauthnRegistration.Challenge,
		UserID:           []byte(account.ID),
		UserVerification: protocol.VerificationRequired,
	}
	credential, err := r.webAuthn.CreateCredential(account, data, parsedResponse)
	if err != nil {
		fmt.Println(err.Error())
		return false, errors.New("WEBAUTHN_REGISTRATION_FAILED")
	}

	cred := m.NewWebauthnCredential()
	cred.PublicKeyID = credential.ID
	cred.PublicKey = credential.PublicKey
	if err := r.accountRepo.CreateWebauthnCredential(ctx, account, cred); err != nil {
		fmt.Println(err.Error())
		return false, errors.New("INAPPROPRIATE_WEBAUTHN_REGISTRATION_CREDENTIAL")
	}

	return true, nil
}
