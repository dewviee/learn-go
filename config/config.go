package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Database DBConfig
	Server   ServerConfig
}

type DBConfig struct {
	NAME     string
	PASSWORD string
	PORT     string
}

type ServerConfig struct {
	Port        int32    `env:"PORT ,decodeunset"`
	AllowOrigin []string `env:"ALLOW_ORIGIN"`
}

// Docs: https://github.com/sethvargo/go-envconfig?tab=readme-ov-file#configuration
func Get(ctx context.Context) (Config, error) {
	var config Config
	err := envconfig.Process(ctx, &config)

	return config, err
}
