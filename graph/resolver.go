package graph

import "graphql/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	beatdrops []*model.Beat
	users []*model.User
	comments []*model.Comment
}
