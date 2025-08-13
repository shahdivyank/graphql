package graph

import (
	"graphql/graph/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	comments []*model.Comment
	db *pgxpool.Pool
}

func NewResolver (db *pgxpool.Pool) *Resolver {
	return &Resolver{
		comments: []*model.Comment{},
		db: db,
	}
}