package config

import (
	"context"
	"log/slog"

	"github.com/caarlos0/env"
)

type Config struct {
	// DB connection info
	DBHost     string `env:"POSTGRES_HOST"`
	DBPort     string `env:"POSTGRES_PORT"`
	DBUser     string `env:"POSTGRES_USER"`
	DBPassword string `env:"POSTGRES_PASSWORD"`
	DBName     string `env:"POSTGRES_DB"`
	DBSslmode  string `env:"POSTGRES_SSLMODE"`

	// Application port number
	Port string `env:"APP_PORT"`
}

func NewConfig(ctx context.Context) (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		slog.ErrorContext(ctx, "failed to load the environment variables", "error", err)
		return nil, err
	}

	return cfg, nil
}
