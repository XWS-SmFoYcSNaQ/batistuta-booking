package startup

import "os"

type Config struct {
	Port         string
	RatingDBHost string
	RatingDBPort string
	RatingDBUser string
	RatingDBPass string
	RatingDBName string
	NatsHost     string
	NatsPort     string
	NatsUser     string
	NatsPass     string
}

func NewConfig() *Config {
	return &Config{
		Port:         os.Getenv("RATING_SERVICE_PORT"),
		RatingDBHost: os.Getenv("RATING_DB_HOST"),
		RatingDBPort: os.Getenv("RATING_DB_PORT"),
		RatingDBUser: os.Getenv("RATING_DB_USER"),
		RatingDBPass: os.Getenv("RATING_DB_PASS"),
		RatingDBName: os.Getenv("RATING_DB_NAME"),
		NatsHost:     os.Getenv("NATS_HOST"),
		NatsPort:     os.Getenv("NATS_PORT"),
		NatsUser:     os.Getenv("NATS_USER"),
		NatsPass:     os.Getenv("NATS_PASS"),
	}
}
