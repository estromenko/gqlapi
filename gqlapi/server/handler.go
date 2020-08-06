package server

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

func (s *Server) handler(schema graphql.Schema) http.HandlerFunc {
	type request struct {
		Query string `json:"query"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		var req request
		json.NewDecoder(r.Body).Decode(&req)

		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: req.Query,
		})

		json.NewEncoder(w).Encode(result)
	}
}
