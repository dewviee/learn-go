package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dewviee/learn-go/config"
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

	fmt.Printf("Server running on port %d\n", cfg.Server.Port)
}
