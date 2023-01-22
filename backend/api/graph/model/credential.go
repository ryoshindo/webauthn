package model

type Credential struct {
	Type                    string                `json:"type"`
	ID                      string                `json:"id"`
	RawID                   string                `json:"rawId"`
	AuthenticatorAttachment string                `json:"authenticatorAttachment"`
	Response                response              `json:"response"`
	ClientExtensionResults  clientExtensionResult `json:"clientExtensionResults"`
}

type response struct {
	ClientDataJson    string   `json:"clientDataJSON"`
	AttestationObject string   `json:"attestationObject"`
	Transports        []string `json:"transports"`
}

type clientExtensionResult struct{}

type CredentialInput struct {
	CompleteWebauthnRegistrationInput
}

func (i *CompleteWebauthnRegistrationInput) Read([]byte) (int, error) {
	return 0, nil
}
