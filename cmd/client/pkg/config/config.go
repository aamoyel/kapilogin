package config

import (
	"fmt"
	"os"
)

type Config struct {
	CapiClusterEndpoint string `json:"capiClusterEndpoint,omitempty" yaml:"capiClusterEndpoint,omitempty"`
	OidcIssuerUrl       string `json:"oidcIssuerUrl,omitempty" yaml:"oidcIssuerUrl,omitempty"`
	OidcClientId        string `json:"oidcClientId,omitempty" yaml:"oidcClientId,omitempty"`
	OidcClientSecret    string `json:"oidcClientSecret" yaml:"oidcClientSecret"`
}

func GetConfig(cfgUri string) (c *Config, err error) {
	envValue, exists := os.LookupEnv("KAPILOGIN_CONFIG")
	if exists {
		cfgUri = envValue
	}

	if cfgUri == "" {
		fmt.Println("Error: KAPILOGIN_CONFIG and config flag is empty\nUse 'kapilogin [command] --help' for more information about a command")
		os.Exit(1)
	}

	fmt.Println(cfgUri)

	return c, nil
}
