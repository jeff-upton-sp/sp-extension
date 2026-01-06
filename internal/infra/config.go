package infra

import (
	"fmt"
	"time"

	"github.com/sailpoint/atlas-go/atlas/config"
)

type Config struct {
	FunctionCacheSize     int
	FunctionCacheDuration time.Duration
}

func loadConfig(s config.Source) (*Config, error) {
	cfg := &Config{}
	cfg.FunctionCacheSize = config.GetInt(s, "FUNCTION_CACHE_SIZE", 4096)
	cfg.FunctionCacheDuration = config.GetDuration(s, "FUNCTION_CACHE_DURATION", 30*time.Second)

	return cfg, cfg.Validate()
}

func (cfg *Config) Validate() error {
	if cfg.FunctionCacheSize < 0 {
		return fmt.Errorf("FunctionCacheSize must be non-negative")
	}

	if cfg.FunctionCacheDuration < 0 {
		return fmt.Errorf("FunctionCacheDuration must be non-negative")
	}

	return nil
}
