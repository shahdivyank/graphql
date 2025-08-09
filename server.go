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
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

type User struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Bio string `json:"bio"`
}

type Beat struct {
	ID uuid.UUID `json:"id"`
	User uuid.UUID `json:"userid"`
	Location string `json:"location"`
	Timestamp int32 `json:"timestamp"`
	Song string `json:"song"`
	Artist string `json:"artist"`
	Description string `json:"description"`
	Longitude string `json:"longitude"`
	Latitude string `json:"latitude"`
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	POSTGRES_URL := os.Getenv("POSTGRES_URL")

	pool, err := pgxpool.New(context.Background(), POSTGRES_URL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	rows, err := pool.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Bio); err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
		users = append(users, user)
	}
	fmt.Println("Users:", users)

	rows, err = pool.Query(context.Background(), "SELECT * FROM beats")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	var beats []Beat
	for rows.Next() {
		var beat Beat
		if err := rows.Scan(&beat.ID, &beat.User, &beat.Timestamp, &beat.Location, &beat.Song, &beat.Artist, &beat.Description, &beat.Longitude, &beat.Latitude); err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
		beats = append(beats, beat)
	}
	fmt.Println("Beats:", beats)

	defer pool.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver()}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
