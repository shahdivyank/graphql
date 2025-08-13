package handler

import (
	database "graphql/db"
	"graphql/gql"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
)


func Handler(w http.ResponseWriter, r *http.Request) {
	connection := database.Connect()
	defer connection.Close()

	srv := gql.Query(connection)

	switch r.URL.Path {
		case "/":
			playground.Handler("GraphQL playground", "/query").ServeHTTP(w, r)
		case "/query":
			srv.ServeHTTP(w, r)
		default: 
			http.NotFound(w, r)
	}
}

 