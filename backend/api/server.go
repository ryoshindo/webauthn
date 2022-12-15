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
	"github.com/ryoshindo/webauthn/backend/api/graph"
	"github.com/ryoshindo/webauthn/backend/api/graph/generated"
	"github.com/ryoshindo/webauthn/backend/db"
	"github.com/ryoshindo/webauthn/backend/repository/psql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("DB_DSN"))))
	db := bun.NewDB(sqldb, pgdialect.New())

	r.Use(dbMiddlewareContext(db))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:*"},
    	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    	ExposedHeaders:   []string{"Link"},
    	AllowCredentials: false,
    	MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: newResolver()}))

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalln(err)
	}
}

func newResolver() *graph.Resolver {
	return graph.NewResolver(
		psql.NewAccountRepository(),
	)
}

func dbMiddlewareContext(d *bun.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = db.Set(ctx, d)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
