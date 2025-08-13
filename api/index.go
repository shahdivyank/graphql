package handler

import (
	database "graphql/db"
	"graphql/gql"
	"net/http"
)


func Handler(w http.ResponseWriter, r *http.Request) {
	connection := database.Connect()
	defer connection.Close()

	srv := gql.Query(connection)

	if r.URL.Path == "/" {
		gql.Playground().ServeHTTP(w, r)
	}
	if r.URL.Path == "/query" {
		srv.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}

 