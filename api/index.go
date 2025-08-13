package handler

import (
	database "graphql/db"
	"graphql/gql"
	"net/http"
)


func Handler(w http.ResponseWriter, r *http.Request) {
	connection := database.Connect()
	defer connection.Close()

	switch r.URL.Path {
		case "/":
			gql.Playground().ServeHTTP(w, r)
		case "/query":
			gql.Query(connection).ServeHTTP(w, r)
		default: 
			http.NotFound(w, r)
	}
}

 