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
	config *config.Config
	logger *zap.Logger
	schema *schema.Schema
}

// NewServer ...
func NewServer(config *config.Config, db *database.Database, logger *zap.Logger, schema *schema.Schema) *Server {
	return &Server{
		db:     db,
		config: config,
		logger: logger,
		schema: schema,
	}
}

// Run ...
func (s *Server) Run() error {

	schema := s.schema.Build()

	http.Handle("/graphql",
		s.baseMiddleware(
			s.authenticationMiddleware(
				s.handler(schema),
			),
		),
	)

	s.logger.Info("Server started at port " + s.config.Server.Port)
	return http.ListenAndServe(":"+s.config.Server.Port, nil)
}
