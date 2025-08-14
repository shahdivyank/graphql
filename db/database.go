package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() *pgxpool.Pool {
	POSTGRES_URL := os.Getenv("POSTGRES_URL")

	connection, err := pgxpool.New(context.Background(), POSTGRES_URL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection: %v\n", err)
		os.Exit(1)
	}

	return connection
}
