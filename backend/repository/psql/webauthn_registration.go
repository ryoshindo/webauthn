package psql

import (
	"context"

	"github.com/ryoshindo/webauthn/backend/db"
	"github.com/ryoshindo/webauthn/backend/model"
)

type WebauthnRegistrationRepository struct{}

func NewWebauthnRegistrationRepository() *WebauthnRegistrationRepository {
	return &WebauthnRegistrationRepository{}
}

func (r *WebauthnRegistrationRepository) FindByID(ctx context.Context, id string) (*model.WebauthnRegistration, error) {
	registration := &model.WebauthnRegistration{}
	if err := db.Get(ctx).NewSelect().Model(registration).Where("id = ?", id).Scan(ctx); err != nil {
		return &model.WebauthnRegistration{}, err
	}

	return registration, nil
}

func (r *WebauthnRegistrationRepository) Create(ctx context.Context, registration *model.WebauthnRegistration) error {
	if _, err := db.Get(ctx).NewInsert().Model(registration).Exec(ctx); err != nil {
		return err
	}

	return nil
}
