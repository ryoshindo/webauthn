package redis

import (
	"context"
	"encoding/json"

	"github.com/ryoshindo/webauthn/backend/kvs"
	"github.com/ryoshindo/webauthn/backend/model"
)

type SessionRepository struct{}

func NewSessionRepository() *SessionRepository {
	return &SessionRepository{}
}

func (r *SessionRepository) FindByToken(ctx context.Context, token string) (*model.Session, error) {
	s, err := kvs.Get(ctx).Get(ctx, token).Result()
	if err != nil {
		return nil, err
	}

	session := model.Session{}
	if err := json.Unmarshal([]byte(s), &session); err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *SessionRepository) Set(ctx context.Context, session *model.Session) error {
	s, err := json.Marshal(session)
	if err != nil {
		return err
	}

	return kvs.Get(ctx).Set(ctx, session.Token, s, session.ExpiresAt).Err()
}

func (r *SessionRepository) Delete(ctx context.Context, session *model.Session) error {
	return kvs.Get(ctx).Del(ctx, session.Token).Err()
}
