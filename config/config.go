package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	Database database
}

type database struct {
	URL string
}

func findRoot() (string, bool) {
	dir, err := os.Getwd()
	if err != nil {
		return "", false
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, true
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", false
		}

		dir = parent
	}
}

func LoadEnv(fileName string) {
	if root, ok := findRoot(); ok {
		if err := godotenv.Load(filepath.Join(root, fileName)); err == nil {
			return
		}
	}

	if err := godotenv.Load(fileName); err != nil {
		log.Printf("warning: could not load %s: %v", fileName, err)
	}
}

func New() *Config {
	godotenv.Load()

	return &Config{
		Database: database{
			URL: os.Getenv("DB_URL"),
		},
	}

}
