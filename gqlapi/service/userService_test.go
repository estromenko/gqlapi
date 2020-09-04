package service

import (
	"gqlapi/config"
	"gqlapi/database"
	"gqlapi/logging"
	"gqlapi/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	conf, err := config.ReadConfig("../config/config_test.yml")
	assert.NoError(t, err, "Error reading config file")

	logger, err := logging.NewLogger(conf)
	assert.NoError(t, err, "Error creating new logger")

	db := database.NewDatabase(conf, logger)

	serv := NewUserService(db, logger, conf)

	tests := []struct {
		name              string
		user              models.User
		validErr          string
		toCompare         string
		passwordsEquility bool
	}{
		{
			name: "Right user",
			user: models.User{
				Email:    "asdasdaf@mail.com",
				Username: "somerandomusername",
				Password: "some valid password",
			},
			validErr:          "",
			toCompare:         "some valid password",
			passwordsEquility: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, serv.validate(&tt.user), tt.validErr)
			assert.NotEmpty(t, serv.hashPassword(tt.user.Password))

			tt.user.Password = serv.hashPassword(tt.user.Password)

			assert.Equal(t, serv.ComparePasswords(&tt.user, tt.toCompare), tt.passwordsEquility)
		})
	}
}
