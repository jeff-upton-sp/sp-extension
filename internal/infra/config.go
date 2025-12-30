package infra

import (
	"github.com/sailpoint/atlas-go/atlas/config"
)

type Config struct {
}

func loadConfig(s config.Source) (*Config, error) {
	cfg := &Config{}
	return cfg, cfg.Validate()
}

func (cfg *Config) Validate() error {
	return nil
}
