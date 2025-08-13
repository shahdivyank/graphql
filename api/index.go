package handler

import (
	"context"
	"fmt"
	"graphql/graph"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vektah/gqlparser/v2/ast"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
	POSTGRES_URL := os.Getenv("POSTGRES_URL")

	connection, err := pgxpool.New(context.Background(), POSTGRES_URL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection: %v\n", err)
		os.Exit(1)
	}

	defer connection.Close()

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver(connection)}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	switch r.URL.Path {
		case "/query":
			srv.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
	}
}