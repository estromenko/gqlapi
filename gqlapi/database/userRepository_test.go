package database

import (
	"gqlapi/config"
	"gqlapi/logging"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	conf, err := config.ReadConfig("../config/config_test.yml")
	assert.NoError(t, err)

	log, err := logging.NewLogger(conf)
	assert.NoError(t, err)

	db := NewDatabase(conf, log)
	assert.NoError(t, db.Open())
	defer db.Close()

	user, err := db.User().FindByID("1")
	assert.NoError(t, err)
	assert.Equal(t, user.ID, 1, "Wrong user id")
}
