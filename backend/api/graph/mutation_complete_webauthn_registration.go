package graph

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ryoshindo/webauthn/backend/api/graph/model"
	"github.com/ryoshindo/webauthn/backend/api/graph/session"
)

func (r *mutationResolver) CompleteWebauthnRegistration(ctx context.Context, input model.CompleteWebauthnRegistrationInput) (bool, error) {
	account := session.AccountFromSession(ctx)
	if account == nil {
		return false, errors.New("UNAUTHORIZED")
	}

	credential := &model.Credential{}
	if err := json.Unmarshal([]byte(input.Credential), credential); err != nil {
		fmt.Println(err.Error())
		return false, errors.New("INAPPROPRIATE_WEBAUTHN_REGISTRATION_CREDENTIAL")
	}

	clientDataBase64Url, err := base64.RawURLEncoding.DecodeString(credential.Response.ClientDataJson)
	if err != nil {
		fmt.Println(err.Error())
		return false, errors.New("INAPPROPRIATE_WEBAUTHN_REGISTRATION_CREDENTIAL_CLIENT_DATA_BASE64URL_DECODING")
	}

	clientData := &model.ClientData{}
	if err := json.Unmarshal(clientDataBase64Url, clientData); err != nil {
		fmt.Println(err.Error())
		return false, errors.New("INAPPROPRIATE_WEBAUTHN_REGISTRATION_CREDENTIAL_CLIENT_DATA_JSON_UNMARSHAL")
	}

	sess := session.GetSession(ctx)
	if sess.Account.WebauthnRegistration.Challenge != clientData.Challenge {
		return false, errors.New("WEBAUTHN_CHALLENGE_NOT_MATCH")
	}

	return true, nil
}
