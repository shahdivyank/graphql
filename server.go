package main

import (
	database "graphql/db"
	"graphql/gql"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)


func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	connection := database.Connect()

	defer connection.Close()

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", gql.Query(connection))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
