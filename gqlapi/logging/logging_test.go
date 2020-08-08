package logging

import (
	"gqlapi/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	conf, err := config.ReadConfig("../config/config_test.yml")
	assert.NoError(t, err)

	logger, err := NewLogger(conf)
	assert.NoError(t, err, "Error creating logger")

	logger.Info("asdasdasdasd")
}
