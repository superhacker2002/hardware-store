package config

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	"log"
)

type Config struct {
	Db string `env:"DATABASE_URL,default=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"`
}

func New() (Config, error) {
	var c Config
	if err := godotenv.Load(); err != nil {
		log.Println("config loading from .env failed:", err)
	}

	ctx := context.Background()
	if err := envconfig.Process(ctx, &c); err != nil {
		return c, err
	}

	return c, nil
}
