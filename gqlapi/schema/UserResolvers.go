package schema

import (
	"gqlapi/database/models"

	"github.com/graphql-go/graphql"
)

func (s *Schema) userResolver(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"].(string)
	user := s.deps.db.User().FindByID(id)
	return user, nil
}

func (s *Schema) usersResolver(p graphql.ResolveParams) (interface{}, error) {
	users, err := s.deps.db.User().FindAll()
	return users, err
}

func (s *Schema) createUserResolver(p graphql.ResolveParams) (interface{}, error) {

	user := models.User{
		Email:    p.Args["email"].(string),
		Username: p.Args["username"].(string),
		Password: p.Args["password"].(string),
	}

	err := s.deps.userService.Create(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
