package server

import (
	"gqlapi/config"
	"gqlapi/schema"
	"net/http"

	"go.uber.org/zap"
)

// Server ...
type Server struct {
	config *config.ServerConfig
	logger *zap.Logger
	schema *schema.Schema
}

// NewServer ...
func NewServer(config *config.Config, logger *zap.Logger, schema *schema.Schema) *Server {
	return &Server{
		config: config.Server,
		logger: logger,
		schema: schema,
	}
}

// Run ...
func (s *Server) Run() error {

	schema := s.schema.Build()

	http.Handle("/graphql", s.baseMiddleware(s.handler(schema)))

	s.logger.Info("Server started at port " + s.config.Port)
	return http.ListenAndServe(":"+s.config.Port, nil)
}
