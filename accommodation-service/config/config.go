package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Address                    string
	AuthServiceAddress         string
	DBAddress                  string
	DBUsername                 string
	DBPassword                 string
	NatsHost                   string
	NatsPort                   string
	NatsUser                   string
	NatsPass                   string
	CreateRatingCommandSubject string
	CreateRatingReplySubject   string
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
		Address:                    os.Getenv("ACCOMMODATION_SERVICE_ADDRESS"),
		AuthServiceAddress:         os.Getenv("AUTH_SERVICE_ADDRESS"),
		DBAddress:                  os.Getenv("ACCOMMODATION_DATABASE_ADDRESS"),
		DBUsername:                 os.Getenv("ACCOMMODATION_DATABASE_USERNAME"),
		DBPassword:                 os.Getenv("ACCOMMODATION_DATABASE_PASSWORD"),
		NatsHost:                   os.Getenv("NATS_HOST"),
		NatsPort:                   os.Getenv("NATS_PORT"),
		NatsUser:                   os.Getenv("NATS_USER"),
		NatsPass:                   os.Getenv("NATS_PASS"),
		CreateRatingCommandSubject: os.Getenv("CREATE_RATING_COMMAND_SUBJECT"),
		CreateRatingReplySubject:   os.Getenv("CREATE_RATING_REPLY_SUBJECT"),
	}
}

func getConfigFromFile() Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}
	return getConfig()
}
