package kvs

import (
	"context"

	"github.com/go-redis/redis/v9"
)

type contextKey struct {
	name string
}

var kvsContextKey = &contextKey{"kvs"}

func Set(ctx context.Context, client *redis.Client) context.Context {
	return context.WithValue(ctx, kvsContextKey, client)
}

func Get(ctx context.Context) *redis.Client {
	if client, ok := ctx.Value(kvsContextKey).(*redis.Client); ok {
		return client
	}
	panic("database connection is not configured in the current context")
}
