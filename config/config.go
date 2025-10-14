package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   dbConfig
	Auth authConfig
	Port string
}

type dbConfig struct {
	Dsn string
}

type authConfig struct {
	Secret string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load env file, using the default one")
	}
	return &Config{
		Db: dbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: authConfig{
			Secret: os.Getenv("SECRET"),
		},
		Port: os.Getenv("PORT"),
	}
}
