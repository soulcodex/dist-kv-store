package config

import (
	_ "github.com/joho/godotenv/autoload" // Autoload env vars from a .env file.
	"github.com/kelseyhightower/envconfig"

	"codesignal/internal/pkg/server"
	"codesignal/internal/pkg/store"
)

// Config contains all the config
// parameters that this service uses.
type (
	Config struct {
		Server     server.Config `envconfig:"SERVER"`
		NodeConfig store.Node
	}
)

// LoadFromEnv will load the env vars from the OS.
func LoadFromEnv() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	return cfg, err
}
