package graph

import (
	"graphql/graph/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	beatdrops map[uuid.UUID]*model.Beat
	users map[uuid.UUID]*model.User
	comments []*model.Comment
	db *pgxpool.Pool
}

func NewResolver (db *pgxpool.Pool) *Resolver {
	return &Resolver{
		beatdrops: make(map[uuid.UUID]*model.Beat),
		comments: []*model.Comment{},
		users: make(map[uuid.UUID]*model.User),
		db: db,
	}
}