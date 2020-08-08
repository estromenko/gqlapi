package database

import (
	"gqlapi/config"
	"gqlapi/logging"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestDatabase(t *testing.T) {
	conf, err := config.ReadConfig("../config/config_test.yml")
	assert.NoError(t, err)

	log, err := logging.NewLogger(conf)
	assert.NoError(t, err)

	tests := []struct {
		name   string
		conf   *config.Config
		logger *zap.Logger
		err    bool
	}{
		{
			name:   "Right config",
			conf:   conf,
			logger: log,
			err:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := NewDatabase(tt.conf, tt.logger)
			err := db.Open()
			assert.Equal(t, err != nil, tt.err, "Error opening database connection")
			if err == nil {
				defer db.Close()
			}
		})
	}
}
