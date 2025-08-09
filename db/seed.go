package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	POSTGRES_URL := os.Getenv("POSTGRES_URL")
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, POSTGRES_URL)

	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

	id := uuid.New()

	_, err = pool.Exec(ctx, `INSERT INTO users (id, name, username, bio) VALUES ($1, $2, $3, $4)`,
		id, "Divyank Shah", "webdiv", "UCR '25")

	if err != nil {
		log.Fatal(err)
	}

	_, err = pool.Exec(ctx, `INSERT INTO beats (id, userid, timestamp, location, song, artist, description, longitude, latitude) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		uuid.New(), id, int32(time.Now().Unix()), "Fremont, CA", "Heaven", "Ailee", "The origins of Beatdrop", 37.12345, 120.12345)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database Seeded Successfully")
}
