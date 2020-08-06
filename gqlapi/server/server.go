package server

import (
	"gqlapi/config"
	"gqlapi/database"
	"gqlapi/schema"
	"net/http"

	"go.uber.org/zap"
)

// Server ...
type Server struct {
	db     *database.Database
	config *config.ServerConfig
	logger *zap.Logger
	schema *schema.Schema
}

// NewServer ...
func NewServer(db *database.Database, config *config.Config, logger *zap.Logger, schema *schema.Schema) *Server {
	return &Server{
		db:     db,
		config: config.Server,
		logger: logger,
		schema: schema,
	}
}

// Run ...
func (s *Server) Run() error {

	schema := s.schema.Build()

	http.Handle("/graphql", s.handler(schema))

	s.logger.Info("Server started at port" + s.config.Port)
	return http.ListenAndServe(":"+s.config.Port, nil)
}
