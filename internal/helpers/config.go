package helpers

import (
	"fmt"

	"github.com/caarlos0/env"
)

// AppConfig - represents application configuration
type AppConfig struct {
	Environment string `env:"ENVIRONMENT" validate:"required"`
	Port        string `env:"PORT" validate:"required"`
	DSN         string `env:"DSN" validate:"required"`

	DisableCaller     bool   `env:"DISABLE_CALLER" validate:"required"`
	DisableStacktrace bool   `env:"DISABLE_STACKTRACE" validate:"required"`
	Encoding          string `env:"ENCODING" validate:"required"`
	Level             string `env:"LEVEL" validate:"required"`

	RedisHost     string `env:"REDIS_HOST" validate:"required"`
	RedisPort     string `env:"REDIS_PORT" validate:"required"`
	RedisPassword string `env:"REDIS_PASSWORD" validate:"required"`
	RedisPoolSize int    `env:"REDIS_POOLSIZE" validate:"required"`
	RedisDB       int    `env:"REDIS_DB" validate:"required"`
}

// LoadConfig - reads configuration from the environment variables.
func LoadConfig() (config AppConfig, _ error) {
	if err := env.Parse(&config); err != nil {
		return config, fmt.Errorf("failed to read configuration: %v", err)
	}
	return config, nil
}
