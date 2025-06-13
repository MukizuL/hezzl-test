package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"go.uber.org/fx"
)

type Config struct {
	Addr     string `env:"RUN_ADDRESS"`
	PgDsn    string `env:"POSTGRES_DSN"`
	ChAddr   string `env:"CLICKHOUSE_ADDRESS"`
	ChUser   string `env:"CLICKHOUSE_USER"`
	ChPass   string `env:"CLICKHOUSE_PASSWORD"`
	NatsAddr string `env:"NATS_ADDR"`
	RedisDsn string `env:"REDIS_DSN"`
}

func newConfig() (*Config, error) {
	cfg := &Config{}

	err := env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("error parsing env variables")
	}

	if cfg.Addr == "" {
		return nil, fmt.Errorf("missing required environment variable RUN_ADDRESS")
	}

	if cfg.PgDsn == "" {
		return nil, fmt.Errorf("missing required environment variable POSTGRES_DSN")
	}

	if cfg.ChAddr == "" {
		return nil, fmt.Errorf("missing required environment variable CLICKHOUSE_ADDRESS")
	}

	if cfg.ChUser == "" {
		return nil, fmt.Errorf("missing required environment variable CLICKHOUSE_USER")
	}

	if cfg.ChPass == "" {
		return nil, fmt.Errorf("missing required environment variable CLICKHOUSE_PASSWORD")
	}

	if cfg.NatsAddr == "" {
		return nil, fmt.Errorf("missing required environment variable NATS_DSN")
	}

	if cfg.RedisDsn == "" {
		return nil, fmt.Errorf("missing required environment variable REDIS_DSN")
	}

	return cfg, nil
}

func Provide() fx.Option {
	return fx.Provide(newConfig)
}
