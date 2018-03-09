package main

import (
	"encoding/json"
	"log"

	"net/http"

	"github.com/graphql-go/graphql"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var schema graphql.Schema

// This Schema
var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "userType",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

// This Query
var userQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"checkuser": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"]
					name := p.Args["name"]

					v, _ := id.(int)
					n, _ := name.(string)

					log.Printf("id user is:%v and name is :%v", v, n)

					data := &user{
						ID:   v,
						Name: n,
					}
					return data, nil
				},
			},
		},
	})

func init() {
	s, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: userQuery,
	})

	if err != nil {
		log.Printf("can not create schema: %v", err)
	}

	schema = s

}

func hGraphql(w http.ResponseWriter, r *http.Request) {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: r.URL.Query().Get("query"),
	})

	if len(result.Errors) > 0 {
		log.Printf("Error create schema: %v", result.Errors)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/graphql", hGraphql)
	http.ListenAndServe(":1234", nil)
	//http://localhost:1234/graphql?query={checkuser(id:1,name:%22ali%22){id,name}}
}
