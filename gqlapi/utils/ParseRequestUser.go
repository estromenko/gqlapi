package utils

import (
	"fmt"

	"github.com/graphql-go/graphql"

	"gqlapi/models"
)

// ParseRequestUser ...
func ParseRequestUser(p graphql.ResolveParams) (*models.User, error) {

	if err := p.Context.Value(ContextValue("tokenError")); err != nil {
		return nil, fmt.Errorf(err.(string))
	}

	if user := p.Context.Value("user"); user != nil {
		if user, ok := user.(models.User); ok {
			return &user, nil
		}
	}

	return nil, fmt.Errorf("Wrong user in context")
}
