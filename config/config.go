package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Database DBConfig
	Server   ServerConfig
}

// Docs: https://github.com/sethvargo/go-envconfig?tab=readme-ov-file#configuration
func Get(ctx context.Context) (Config, error) {
	var config Config
	err := envconfig.Process(ctx, &config)

	return config, err
}
