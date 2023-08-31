package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	const configVarName string = "KAPILOGIN_CONFIG"

	for _, test := range []struct {
		description        string
		configVar          string
		configVarValue     string
		configFlagValue    string
		shouldInjectEnvVar bool
	}{
		{
			description:        "Should return error when env var and flag are not set",
			configVar:          configVarName,
			configVarValue:     "",
			shouldInjectEnvVar: false,
		},
	} {
		t.Run(test.description, func(t *testing.T) {
			_, err := GetConfig("")
			assert.NotNil(t, err)
		})
	}
}
