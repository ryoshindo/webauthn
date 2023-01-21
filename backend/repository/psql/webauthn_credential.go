package psql

import (
	"context"

	"github.com/ryoshindo/webauthn/backend/db"
	"github.com/ryoshindo/webauthn/backend/model"
)

type WebauthnCredentialRepository struct{}

func NewWebauthnCredentialRepository() *WebauthnCredentialRepository {
	return &WebauthnCredentialRepository{}
}

func (r *WebauthnCredentialRepository) FindByID(ctx context.Context, id string) (*model.WebauthnCredential, error) {
	credential := &model.WebauthnCredential{}
	if err := db.Get(ctx).NewSelect().Model(credential).Where("id = ?", id).Scan(ctx); err != nil {
		return &model.WebauthnCredential{}, err
	}

	return credential, nil
}

func (r *WebauthnCredentialRepository) Create(ctx context.Context, credential *model.WebauthnCredential) error {
	if _, err := db.Get(ctx).NewInsert().Model(credential).Exec(ctx); err != nil {
		return err
	}

	return nil
}
