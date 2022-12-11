package db

import (
	"context"

	"github.com/uptrace/bun"
)

type contextKey struct {
	name string
}

var dbKey = &contextKey{"db"}

func Set(ctx context.Context, db *bun.DB) context.Context {
	return context.WithValue(ctx, dbKey, db)
}

func Get(ctx context.Context) *bun.DB {
	if db, ok := ctx.Value(dbKey).(*bun.DB); ok {
		return db
	}
	panic("database connection is not configured in the current context")
}
