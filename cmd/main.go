package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dewviee/learn-go/config"
	"github.com/dewviee/learn-go/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	ctx := context.Background()

	if os.Getenv("TYPE") == "development" {
		config.LoadEnv()
	}
	cfg, err := config.Get(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	_ = database.Init(cfg.Database)

	app := fiber.New(config.GetFiberConfig(cfg.Server))

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	log.Fatal(app.Listen(fmt.Sprintf("%s:%d", cfg.Server.Address, cfg.Server.Port)))
}
