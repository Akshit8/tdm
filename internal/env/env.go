// Package env provides methods and types to pick secrets and configuration at run-time.
package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o envtesting/provider.gen.go . Provider

// Provider defines methods for external secret provider like Hashicorp vault.
type Provider interface {
	Get(key string) (string, error)
}

// Configuration provides object to access secrets from env file and provider
type Configuration struct {
	provider Provider
}

func NewConfiguration(provider Provider) *Configuration {
	return &Configuration{
		provider: provider,
	}
}

// Load read the env filename and load it into ENV for this process.
func Load(filepath string) error {
	err := godotenv.Load(filepath)
	if err != nil {
		return fmt.Errorf("loading env file: %w", err)
	}

	return nil
}

// Get returns the value from environment variable `<key>`. When an environment variable `<key>_SECURE` exists
// the provider is used for getting the value.
func (c *Configuration) Get(key string) (string, error) {
	val := os.Getenv(key)
	secretKey := os.Getenv(fmt.Sprintf("%s_SECURE", key))

	if secretKey != "" {
		secretKeyVal, err := c.provider.Get(secretKey)
		if err != nil {
			return "", fmt.Errorf("provider get: %w", err)
		}

		val = secretKeyVal
	}

	return val, nil
}
