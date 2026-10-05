package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database database
}

type database struct {
	URL string
}

func New() *Config {
	godotenv.Load()

	return &Config{
		Database: database{
			URL: os.Getenv("DATABASE_URL"),
		},
	}

}
