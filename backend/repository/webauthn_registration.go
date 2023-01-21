package repository

import (
	"context"

	"github.com/ryoshindo/webauthn/backend/model"
)

type WebauthnRegistrationRepository interface {
	FindByID(context.Context, string) (*model.WebauthnRegistration, error)
	Create(context.Context, *model.WebauthnRegistration) error
}
