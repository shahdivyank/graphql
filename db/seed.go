package database

import (
	"context"
	"encoding/json"
	"graphql/graph/model"
	"io"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func insert_users(db *pgxpool.Pool) {
	file, err := os.Open("db/users.json")

	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var users []model.User

	if err := json.Unmarshal(data, &users); err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		_, err = db.Exec(context.Background(), `INSERT INTO users (id, name, username, bio, beatdrops, friends, settings, photo) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
			user.ID, user.Name, user.Username, user.Bio, user.Beatdrops, user.Friends, "{}", user.Photo)

		if err != nil {
			log.Fatalf("Failed to insert user: %v", err)
		}
	}
}

type Comment struct {
	ID string
	Userid string
	Beatid string
	Timestamp string
	Comment string
}

func insert_comments(db *pgxpool.Pool) {
	file, err := os.Open("db/comments.json")

	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var comments []Comment

	if err := json.Unmarshal(data, &comments); err != nil {
		log.Fatal(err)
	}

	for _, comment := range comments {

		_, err = db.Exec(context.Background(), `INSERT INTO comments (id, userid, beatid, timestamp, comment) VALUES ($1, $2, $3, $4, $5)`,
			comment.ID, comment.Userid, comment.Beatid, comment.Timestamp, comment.Comment)

		if err != nil {
			log.Fatalf("Failed to insert comments: %v", err)
		}
	}
}

type Beat struct  {
	ID string
	Userid string
	Timestamp string
	Location string
	Song string
	Artist string
	Description string
	Longitude float32
	Latitude float32
	Image string
	Comments int32
}

func insert_beats(db *pgxpool.Pool) {
	file, err := os.Open("db/beats.json")

	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var beats []Beat

	if err := json.Unmarshal(data, &beats); err != nil {
		log.Fatal(err)
	}

	for _, beat := range beats {
		_, err = db.Exec(context.Background(), `INSERT INTO beats (id, userid, song, artist, description, location, longitude, latitude, image, timestamp, comments) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
			beat.ID, beat.Userid, beat.Song, beat.Artist, beat.Description, beat.Location, beat.Longitude, beat.Latitude, beat.Image, beat.Timestamp, beat.Comments)

		if err != nil {
			log.Fatalf("Failed to insert user: %v", err)
		}
	}
}

func Seed(db *pgxpool.Pool) {
	insert_users(db)
	insert_beats(db)
	insert_comments(db)
}
