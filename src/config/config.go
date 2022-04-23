package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
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
