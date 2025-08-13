package main

import (
	"context"
	"fmt"
	"graphql/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/ast"
)

var connection *pgxpool.Pool

func database() {
	POSTGRES_URL := os.Getenv("POSTGRES_URL")

	var err error
	connection, err = pgxpool.New(context.Background(), POSTGRES_URL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection: %v\n", err)
		os.Exit(1)
	}
}

func graphql() http.Handler {
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver(connection)}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})
	return srv
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/query" {
		graphql().ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}


func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}

	database()
	defer connection.Close()

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", graphql())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

 
// func Handler(w http.ResponseWriter, r *http.Request) {
// 	POSTGRES_URL := os.Getenv("POSTGRES_URL")

// 	connection, err := pgxpool.New(context.Background(), POSTGRES_URL)

// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to create connection: %v\n", err)
// 		os.Exit(1)
// 	}

// 	defer connection.Close()

// 	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver(connection)}))

// 	srv.AddTransport(transport.Options{})
// 	srv.AddTransport(transport.GET{})
// 	srv.AddTransport(transport.POST{})

// 	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

// 	srv.Use(extension.Introspection{})
// 	srv.Use(extension.AutomaticPersistedQuery{
// 		Cache: lru.New[string](100),
// 	})

// 	switch r.URL.Path {
// 		case "/query":
// 			srv.ServeHTTP(w, r)
// 		default:
// 			http.NotFound(w, r)
// 	}
// }