package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Address                     string
	AuthServiceAddress          string
	AccommodationServiceAddress string
	DBAddress                   string
	DBUsername                  string
	DBPassword                  string
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
		Address:                     os.Getenv("BOOKING_SERVICE_ADDRESS"),
		AuthServiceAddress:          os.Getenv("AUTH_SERVICE_ADDRESS"),
		AccommodationServiceAddress: os.Getenv("ACCOMMODATION_SERVICE_ADDRESS"),
		DBAddress:                   os.Getenv("BOOKING_DATABASE_ADDRESS"),
		DBUsername:                  os.Getenv("BOOKING_DATABASE_USERNAME"),
		DBPassword:                  os.Getenv("BOOKING_DATABASE_PASSWORD"),
	}
}

func getConfigFromFile() Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}
	return getConfig()
}
