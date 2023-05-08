package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Address    string
	DBAddress  string
	DBUsername string
	DBPassword string
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
		Address:    os.Getenv("ACCOMMODATION_SERVICE_ADDRESS"),
		DBAddress:  os.Getenv("ACCOMMODATION_DATABASE_ADDRESS"),
		DBUsername: os.Getenv("ACCOMMODATION_DATABASE_USERNAME"),
		DBPassword: os.Getenv("ACCOMMODATION_DATABASE_PASSWORD"),
	}
}

func getConfigFromFile() Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}
	return getConfig()
}
