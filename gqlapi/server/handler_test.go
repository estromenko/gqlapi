package server

import (
	"encoding/json"
	"gqlapi/config"
	"gqlapi/database"
	"gqlapi/logging"
	"gqlapi/schema"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	conf, err := config.ReadConfig("../config/config_test.yml")
	assert.NoError(t, err, "Error reading config file")
	logger, err := logging.NewLogger(conf)
	assert.NoError(t, err, "Error creating logger with given config")
	db := database.NewDatabase(conf, logger)
	sch := schema.NewSchema(db, logger, conf)
	s := NewServer(conf, logger, sch)

	query := "query{users{email,password}}"
	req, err := http.NewRequest("GET", "/graphql", nil)
	req.URL.Query().Add("query", query)

	r := httptest.NewRecorder()
	http.HandlerFunc(s.handler(s.schema.Build())).ServeHTTP(r, req)

	type response struct {
		Data   map[string]string `json:"data"`
		Errors map[string]string `json:"errors"`
	}
	var res response

	json.NewDecoder(r.Body).Decode(&res)

	assert.NotNil(t, res.Data)

	b, _ := json.Marshal(res)
	logger.Info(string(b))

	assert.Empty(t, res.Errors)
}
