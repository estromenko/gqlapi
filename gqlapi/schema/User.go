package schema

import (
	"github.com/graphql-go/graphql"
)

// UserType ...
func UserType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{

			Name: "User",

			Fields: graphql.Fields{

				"id": &graphql.Field{
					Type: graphql.ID,
				},

				"email": &graphql.Field{
					Type: graphql.String,
				},

				"username": &graphql.Field{
					Type: graphql.String,
				},

				"password": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}