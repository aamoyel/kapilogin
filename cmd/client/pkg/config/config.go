package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	KapiloginApiEndpoint string `json:"kapiloginApiEndpoint,omitempty" yaml:"kapiloginApiEndpoint,omitempty"`
	OidcIssuerUrl        string `json:"oidcIssuerUrl,omitempty" yaml:"oidcIssuerUrl,omitempty"`
	OidcClientId         string `json:"oidcClientId,omitempty" yaml:"oidcClientId,omitempty"`
	OidcClientSecret     string `json:"oidcClientSecret" yaml:"oidcClientSecret"`
}

func GetConfig(cfgUri string) (*Config, error) {
	envValue, exists := os.LookupEnv("KAPILOGIN_CONFIG")
	if exists {
		cfgUri = envValue
	}

	if cfgUri == "" {
		return nil, errors.New("Error: KAPILOGIN_CONFIG and config flag is empty. Use 'kapilogin [command] --help' for more information about a command\n")
	}

	file, err := os.Open(cfgUri)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileAsBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var (
		config    = &Config{}
		extension = filepath.Ext(cfgUri)
	)

	switch extension {
	case ".json":
		if err := json.Unmarshal(fileAsBytes, config); err != nil {
			return nil, fmt.Errorf("failed to create server configuration from file %s: %w", cfgUri, err)
		}
	case ".yml", ".yaml":
		if err := yaml.Unmarshal(fileAsBytes, config); err != nil {
			return nil, fmt.Errorf("failed to create server configuration from file %s: %w", cfgUri, err)
		}
	default:
		return nil, errors.New("invalid server configuration file extension, supported: .json|.yml|.yaml")
	}

	c := config
	return c, nil
}
