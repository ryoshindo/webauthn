package psql

import (
	"context"

	"github.com/ryoshindo/webauthn/backend/db"
	"github.com/ryoshindo/webauthn/backend/model"
)

type AccountRepository struct{}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (r *AccountRepository) FindByID(ctx context.Context, id string) (*model.Account, error) {
	account := &model.Account{}
	if err := db.Get(ctx).NewSelect().Model(account).Where("id = ?", id).Scan(ctx); err != nil {
		return &model.Account{}, err
	}

	return account, nil
}

func (r *AccountRepository) FindByEmail(ctx context.Context, email string) (*model.Account, error) {
	account := &model.Account{}
	if err := db.Get(ctx).NewSelect().Model(account).Where("email = ?", email).Scan(ctx); err != nil {
		return &model.Account{}, err
	}

	return account, nil
}

func (r *AccountRepository) Create(ctx context.Context, account *model.Account) error {
	if _, err := db.Get(ctx).NewInsert().Model(account).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r *AccountRepository) CreateWebauthnCredential(ctx context.Context, account *model.Account, webauthnCredential *model.WebauthnCredential) error {
	webauthnCredential.AccountID = account.ID
	if _, err := db.Get(ctx).NewInsert().Model(webauthnCredential).Exec(ctx); err != nil {
		return err
	}

	return nil
}
