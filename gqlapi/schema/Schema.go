package schema

import (
	"gqlapi/database"

	"github.com/graphql-go/graphql"
	"go.uber.org/zap"
)

// Schema ...
type Schema struct {
	deps *dependencies
}

// NewSchema ...
func NewSchema(db *database.Database, logger *zap.Logger) *Schema {
	return &Schema{
		deps: &dependencies{
			db:     db,
			logger: logger,
		},
	}
}

// Build ...
func (s *Schema) Build() graphql.Schema {
	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: s.QueryType(),
		},
	)
	return schema
}

type dependencies struct {
	db     *database.Database
	logger *zap.Logger
}
