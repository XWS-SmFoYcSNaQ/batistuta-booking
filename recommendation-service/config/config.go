package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Address                     string
	DBAddress                   string
	DBUsername                  string
	DBPassword                  string
	AccommodationServiceAddress string
	BookingServiceAddress       string
	UserServiceAddress          string
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
		Address:                     os.Getenv("RECOMMENDATION_SERVICE_ADDRESS"),
		DBAddress:                   os.Getenv("RECOMMENDATION_DATABASE_ADDRESS"),
		DBUsername:                  os.Getenv("RECOMMENDATION_DATABASE_USERNAME"),
		DBPassword:                  os.Getenv("RECOMMENDATION_DATABASE_PASSWORD"),
		AccommodationServiceAddress: os.Getenv("ACCOMMODATION_SERVICE_ADDRESS"),
		BookingServiceAddress:       os.Getenv("BOOKING_SERVICE_ADDRESS"),
		UserServiceAddress:          os.Getenv("USER_SERVICE_ADDRESS"),
	}
}

func getConfigFromFile() Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}
	return getConfig()
}
