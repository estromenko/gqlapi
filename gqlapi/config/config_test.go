package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	tests := []struct{
		name string
		path string
		err bool
	}{
		{
			name: "Wrong config path",
			path: "Some random config path",
			err: true,
		},
		{
			name: "Right config path",
			path: "config.yml",
			err: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ReadConfig(tt.path)
			assert.Equal(t, err != nil, tt.err, "Unexpected error or no error")
		})
	}
}