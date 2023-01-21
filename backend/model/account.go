package model

import (
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
)

type Account struct {
	ID        string    `bun:"id,pk"`
	CreatedAt time.Time `bun:"created_at"`
	UpdatedAt time.Time `bun:"updated_at"`
	Email     string    `bun:"email"`
	UserName  string    `bun:"user_name"`

	WebauthnRegistrations []*WebauthnRegistration `bun:"rel:has-many,join:id=account_id"`
	WebauthnCredentials   []*WebauthnCredential   `bun:"rel:has-many,join:id=account_id"`

	credential []webauthn.Credential
}

type AccountList []Account

func NewAccount() *Account {
	return &Account{
		ID: NewULIDString(),
	}
}

func (a *Account) AddCredential(cred webauthn.Credential) {
	a.credential = append(a.credential, cred)
}

func (a *Account) CredentialExcludeList() []protocol.CredentialDescriptor {
	credentialExcludeList := []protocol.CredentialDescriptor{}
	for _, cred := range a.credential {
		descriptor := protocol.CredentialDescriptor{
			Type:         protocol.PublicKeyCredentialType,
			CredentialID: cred.ID,
		}
		credentialExcludeList = append(credentialExcludeList, descriptor)
	}

	return credentialExcludeList
}

func (a *Account) WebAuthnID() []byte {
	return []byte(a.ID)
}

func (a *Account) WebAuthnName() string {
	return a.Email
}

func (a *Account) WebAuthnDisplayName() string {
	return a.UserName
}

func (a *Account) WebAuthnIcon() string {
	return ""
}

func (a *Account) WebAuthnCredentials() []webauthn.Credential {
	return a.credential
}
