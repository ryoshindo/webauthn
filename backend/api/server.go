package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-redis/redis/v9"
	"github.com/ryoshindo/webauthn/backend/api/graph"
	"github.com/ryoshindo/webauthn/backend/api/graph/generated"
	"github.com/ryoshindo/webauthn/backend/api/graph/session"
	"github.com/ryoshindo/webauthn/backend/db"
	"github.com/ryoshindo/webauthn/backend/kvs"
	prepo "github.com/ryoshindo/webauthn/backend/repository/psql"
	rrepo "github.com/ryoshindo/webauthn/backend/repository/redis"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := chi.NewRouter()
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3333", "http://localhost:8080"}, // FIXME
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	r.Use(middleware.Logger)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("DB_DSN"))))
	db := bun.NewDB(sqldb, pgdialect.New())
	defer db.Close()
	db.AddQueryHook(bundebug.NewQueryHook())

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // FIXME: hostname
		Password: "",           // no password set
		DB:       0,            // use default DB
	})
	defer redisClient.Close()

	r.Use(dbContextMiddleware(db))
	r.Use(kvsContextMiddleware(redisClient))
	r.Use(newSessionMiddleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: newResolver()}))

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port) // FIXME: hostname
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalln(err)
	}
}

func newResolver() *graph.Resolver {
	return graph.NewResolver(
		prepo.NewAccountRepository(),
		prepo.NewWebauthnCredentialRepository(),
	)
}

func newSessionMiddleware() func(http.Handler) http.Handler {
	return session.Middleware(prepo.NewAccountRepository(), rrepo.NewSessionRepository())
}

func dbContextMiddleware(d *bun.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = db.Set(ctx, d)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func kvsContextMiddleware(c *redis.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = kvs.Set(ctx, c)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
