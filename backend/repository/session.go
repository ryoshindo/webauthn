package repository

import (
	"context"

	"github.com/ryoshindo/webauthn/backend/model"
)

type SessionRepository interface {
	FindByToken(context.Context, string) (*model.Session, error)
	Create(context.Context, *model.Session) error
	Delete(context.Context, *model.Session) error
}
