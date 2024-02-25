package main

import (
	"errors"
	"fmt"
	"os"
)

type Config struct {
	IdentityUrl string
	Environment string
}

func NewConfig() (*Config, error) {
	config := &Config{}

	environment := os.Getenv("environment")
	if environment == "" {
		environment = "prod"
	}
	config.Environment = environment

	switch environment {
	case "prod", "production":
		config.IdentityUrl = "https://identityservice.audius.co"
		break
	case "stage", "staging":
		config.IdentityUrl = "https://identityservice.staging.audius.co"
		break
	case "development", "dev", "local":
		config.IdentityUrl = "http://audius-protocol-identity-service-1"
		break
	default:
		return nil, errors.New(fmt.Sprintf("environment %s not valid", environment))
	}

	logger.Info("configuration", "env", environment, "identity", config.IdentityUrl)

	return config, nil
}
