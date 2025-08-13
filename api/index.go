package handler

import (
	database "graphql/db"
	"graphql/gql"
	"net/http"
)


func Handler(w http.ResponseWriter, r *http.Request) {

	connection := database.Connect()

	if r.URL.Path == "/" {
		gql.Playground().ServeHTTP(w, r)
	}
	if r.URL.Path == "/query" {
		gql.Query(connection).ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}

 