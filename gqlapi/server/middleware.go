package server

import (
	"net/http"
	"time"
)

func (s *Server) baseMiddleware(endpoint http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		s.logger.Info(r.UserAgent() + ". " + r.Method + ". " + time.Now().String())
		endpoint(w, r)
	}
}
