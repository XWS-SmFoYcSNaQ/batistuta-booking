package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Address string
}

func LoadConfig() Config {
	cfg := getConfig()
	if cfg.Address == "" {
		return getConfigFromFile()
	}
	return cfg
}

func getConfig() Config {
	return Config{
		Address: os.Getenv("ACCOMMODATION_SERVICE_ADDRESS"),
	}
}

func getConfigFromFile() Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}
	return getConfig()
}
