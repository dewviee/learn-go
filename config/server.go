package config

import "github.com/gofiber/fiber/v2"

type ServerConfig struct {
	Address     string   `env:"SERVER_HOST ,default="`
	Port        int32    `env:"SERVER_PORT ,decodeunset"`
	AllowOrigin []string `env:"ALLOW_ORIGIN"`
}

func GetFiberConfig(config ServerConfig) fiber.Config {
	return fiber.Config{
		AppName: "Just Post v0.0.1",
	}
}
