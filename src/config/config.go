package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Cannot load .env file")
	}
	return &Config{}
}

func (*Config) Getenv(key string, defVal string) string {
	env := os.Getenv(key)
	if env == "" {
		env = defVal
	}
	return env
}
