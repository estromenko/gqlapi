package schema

import (
	"github.com/graphql-go/graphql"
)

// QueryType ...
func (s *Schema) QueryType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{

			Name: "Query",

			Fields: graphql.Fields{

				"user": &graphql.Field{
					Type: UserType(),
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.ID,
						},
					},
					Resolve: s.userResolver,
				},

				"users": &graphql.Field{
					Type: graphql.NewList(UserType()),
					Resolve: s.usersResolver,
				},
			},
		},
	)
}
