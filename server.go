package main

import (
	"context"
	"fmt"
	database "graphql/db"
	"graphql/gql"
	"graphql/graph/model"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	connection := database.Connect()
	defer connection.Close()

	rows, error := connection.Query(context.Background(), "SELECT id FROM users")

	if error != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", error)
	}

	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User

		if err := rows.Scan(&user.ID); err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}

		users = append(users, &user)
	}

	if len(users) == 0 {
		database.Seed(connection)
	}

	http.Handle("/", gql.Playground())
	http.Handle("/query", gql.Query(connection))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
