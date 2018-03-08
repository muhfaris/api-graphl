package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Schema graphql.Schema

var whoType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "whoami",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"who": &graphql.Field{
				Type: whoType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return p.Context.Value("currentUser"), nil

				},
			},
		},
	})

func init() {
	s, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})

	if err != nil {
		log.Fatalf("Error create schema: %v", err)
	}

	Schema = s
}

func main() {
	http.HandleFunc("/graphql", handleGraphql)
	http.ListenAndServe(":7171", nil)

}

func handleGraphql(w http.ResponseWriter, h *http.Request) {
	u := &user{
		ID:   1,
		Name: "im linuxer",
	}
	result := graphql.Do(graphql.Params{
		Schema:        Schema,
		RequestString: h.URL.Query().Get("query"),
		Context:       context.WithValue(context.Background(), "currentUser", u),
	})

	json.NewEncoder(w).Encode(result)

}
