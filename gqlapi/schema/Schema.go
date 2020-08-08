package schema

import (
	"gqlapi/config"
	"gqlapi/database"
	"gqlapi/service"

	"github.com/graphql-go/graphql"
	"go.uber.org/zap"
)

// Schema ...
type Schema struct {
	deps *dependencies
}

// NewSchema ...
func NewSchema(db *database.Database, logger *zap.Logger, config *config.Config) *Schema {
	return &Schema{
		deps: &dependencies{
			db:          db,
			logger:      logger,
			userService: service.NewUserService(db, logger, config),
		},
	}
}

// Build ...
func (s *Schema) Build() graphql.Schema {
	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    s.QueryType(),
			Mutation: s.MutationType(),
		},
	)
	return schema
}

type dependencies struct {
	db          *database.Database
	logger      *zap.Logger
	userService *service.UserService
}
