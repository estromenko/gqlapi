package server

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

func (s *Server) handler(schema graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		query := parseQuery(r)

		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  query.Query,
			VariableValues: query.Variables,
			OperationName:  query.OperationName,
		})

		json.NewEncoder(w).Encode(result)
	}
}

type request struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

func parseQuery(r *http.Request) request {
	var req request

	if r.Method == "GET" {
		params := r.URL.Query()
		req.Query = params.Get("query")
		req.OperationName = params.Get("operationName")
	} else {
		json.NewDecoder(r.Body).Decode(&req)
	}

	return req
}
