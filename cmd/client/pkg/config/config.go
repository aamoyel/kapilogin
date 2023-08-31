package config

import (
	"errors"
	"os"
)

type Config struct {
	CapiClusterEndpoint string `json:"capiClusterEndpoint,omitempty" yaml:"capiClusterEndpoint,omitempty"`
	OidcIssuerUrl       string `json:"oidcIssuerUrl,omitempty" yaml:"oidcIssuerUrl,omitempty"`
	OidcClientId        string `json:"oidcClientId,omitempty" yaml:"oidcClientId,omitempty"`
	OidcClientSecret    string `json:"oidcClientSecret" yaml:"oidcClientSecret"`
}

func GetConfig(cfgUri string) (*Config, error) {
	envValue, exists := os.LookupEnv("KAPILOGIN_CONFIG")
	if exists {
		cfgUri = envValue
	}

	if cfgUri == "" {
		return nil, errors.New("Error: KAPILOGIN_CONFIG and config flag is empty\nUse 'kapilogin [command] --help' for more information about a command")
	}

	c := &Config{}
	return c, nil
}
