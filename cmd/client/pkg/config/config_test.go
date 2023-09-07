package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	for _, test := range []struct {
		description        string
		config Config
	}{
		{
			description:        "Should return error when required field is empty or not set",
			// TODO test config
		},
	} {
		t.Run(test.description, func(t *testing.T) {
			_, err := GetConfig("")
			assert.NotNil(t, err)
		})
	}
}
