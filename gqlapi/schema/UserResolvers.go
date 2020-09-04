package schema

import (
	"gqlapi/models"
	//	"gqlapi/utils"

	"github.com/graphql-go/graphql"
)

func (s *Schema) userResolver(p graphql.ResolveParams) (interface{}, error) {

	// _, err := utils.ParseRequestUser(p)
	// if err != nil {
	// 	return nil, err
	// }

	id := p.Args["id"].(string)
	user, err := s.deps.db.User().FindByID(id)
	return user, err
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

	token, err := s.deps.userService.Create(&user)
	if err != nil {
		return nil, err
	}

	userWithToken := models.UserWithToken{
		User:  user,
		Token: token,
	}

	return userWithToken, nil
}
