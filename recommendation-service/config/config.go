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
		Address:    os.Getenv("RECOMMENDATION_SERVICE_ADDRESS"),
		DBAddress:  os.Getenv("RECOMMENDATION_DATABASE_ADDRESS"),
		DBUsername: os.Getenv("RECOMMENDATION_DATABASE_USERNAME"),
		DBPassword: os.Getenv("RECOMMENDATION_DATABASE_PASSWORD"),
	}
}

func getConfigFromFile() Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}
	return getConfig()
}
