package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port               int    `envconfig:"PORT" default:"8080"`
	DatabaseURL        string `envconfig:"DATABASE_URL" required:"true"`
	DatabaseMigrations string `envconfig:"DATABASE_MIGRATIONS" default:"file:///migrations"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
