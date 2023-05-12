package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Address                     string
	AccommodationServiceAddress string
	BookingServiceAddress       string
	UserServiceAddress          string
	AuthServiceAddress          string
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
		AccommodationServiceAddress: os.Getenv("ACCOMMODATION_SERVICE_ADDRESS"),
		Address:                     os.Getenv("GATEWAY_ADDRESS"),
		BookingServiceAddress:       os.Getenv("BOOKING_SERVICE_ADDRESS"),
		UserServiceAddress:          os.Getenv("USER_SERVICE_ADDRESS"),
		AuthServiceAddress:          os.Getenv("AUTH_SERVICE_ADDRESS"),
	}
}

func getConfigFromFile() Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}
	return getConfig()
}
