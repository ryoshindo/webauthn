package session

import (
	"context"
	"net/http"

	"github.com/ryoshindo/webauthn/backend/model"
	"github.com/ryoshindo/webauthn/backend/repository"
)

var accountCtxKey = &contextKey{"account"}

type contextKey struct {
	name string
}

func Middleware(accountRepo repository.AccountRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			accountID, ok := ctx.Value(accountCtxKey).(string)
			if !ok {
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			account, err := accountRepo.FindByID(r.Context(), accountID)
			if err != nil {
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			ctx = Set(ctx, account)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Account(ctx context.Context) *model.Account {
	account, _ := ctx.Value(accountCtxKey).(*model.Account)
	return account
}

func Set(ctx context.Context, account *model.Account) context.Context {
	return context.WithValue(ctx, accountCtxKey, account)
}

func Remove(ctx context.Context) context.Context {
	return context.WithValue(ctx, accountCtxKey, nil)
}
