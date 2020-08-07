package schema

import "github.com/graphql-go/graphql"

func (s *Schema) userResolver(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"].(string)
	user := s.deps.db.User().FindByID(id)
	return user, nil
}

func (s *Schema) usersResolver(p graphql.ResolveParams) (interface{}, error) {
	users, err := s.deps.db.User().FindAll()
	return users, err
}
