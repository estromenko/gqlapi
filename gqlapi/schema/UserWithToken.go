package schema

import "github.com/graphql-go/graphql"

// UserWithTokenType ...
func UserWithTokenType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "UserWithToken",
			Fields: graphql.Fields{
				"user": &graphql.Field{
					Type: UserType(),
				},
				"token": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}
