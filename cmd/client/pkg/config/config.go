package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

const _configVarName string = "KAPILOGIN_CONFIG"

type Config struct {
	KapiloginApiEndpoint string `json:"kapiloginApiEndpoint,omitempty" yaml:"kapiloginApiEndpoint,omitempty"`
	OidcIssuerUrl        string `json:"oidcIssuerUrl,omitempty" yaml:"oidcIssuerUrl,omitempty"`
	OidcClientId         string `json:"oidcClientId,omitempty" yaml:"oidcClientId,omitempty"`
	OidcClientSecret     string `json:"oidcClientSecret" yaml:"oidcClientSecret"`
}

func GetConfig(cfgUri string) (*Config, error) {
	envValue, exists := os.LookupEnv(_configVarName)
	if exists {
		cfgUri = envValue
	}

	if cfgUri == "" {
		return nil, errors.New("Error: " + _configVarName + " and config flag is empty. Use 'kapilogin [command] --help' for more information about a command\n")
	}

	var fileAsBytes []byte
	if strings.HasPrefix(cfgUri, "http") {
		response, err := http.Get(cfgUri)
		if err != nil {
			return nil, err
		}

		fileAsBytes, err = io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
	} else {
		file, err := os.Open(cfgUri)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		fileAsBytes, err = io.ReadAll(file)
		if err != nil {
			return nil, err
		}
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

	if config.KapiloginApiEndpoint == "" || config.OidcClientId == "" || config.OidcIssuerUrl == "" {
		return nil, fmt.Errorf("configuration file is invalid, required fields is not set, empy or misspelled %+v", config)
	}

	c := config
	return c, nil
}
