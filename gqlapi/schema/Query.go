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
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						id := p.Args["id"].(string)
						user := s.deps.db.User().FindByID(id)
						return user, nil
					},
				},

				"users": &graphql.Field{
					Type: graphql.NewList(UserType()),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						users, err := s.deps.db.User().FindAll()
						return users, err
					},
				},
			},
		},
	)
}
