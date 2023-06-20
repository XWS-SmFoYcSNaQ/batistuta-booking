package config

import "os"

type Config struct {
	Port                       string
	RatingDBHost               string
	RatingDBPort               string
	RatingDBUser               string
	RatingDBPass               string
	RatingDBName               string
	NatsHost                   string
	NatsPort                   string
	NatsUser                   string
	NatsPass                   string
	AuthServiceAddress         string
	UserServiceAddress         string
	CreateRatingCommandSubject string
	CreateRatingReplySubject   string
	DeleteRatingCommandSubject string
	DeleteRatingReplySubject   string
	NotificationSubject        string
}

func NewConfig() *Config {
	return &Config{
		Port:                       os.Getenv("RATING_SERVICE_PORT"),
		RatingDBHost:               os.Getenv("RATING_DB_HOST"),
		RatingDBPort:               os.Getenv("RATING_DB_PORT"),
		RatingDBUser:               os.Getenv("RATING_DB_USER"),
		RatingDBPass:               os.Getenv("RATING_DB_PASS"),
		RatingDBName:               os.Getenv("RATING_DB_NAME"),
		NatsHost:                   os.Getenv("NATS_HOST"),
		NatsPort:                   os.Getenv("NATS_PORT"),
		NatsUser:                   os.Getenv("NATS_USER"),
		NatsPass:                   os.Getenv("NATS_PASS"),
		AuthServiceAddress:         os.Getenv("AUTH_SERVICE_ADDRESS"),
		UserServiceAddress:         os.Getenv("USER_SERVICE_ADDRESS"),
		CreateRatingCommandSubject: os.Getenv("CREATE_RATING_COMMAND_SUBJECT"),
		CreateRatingReplySubject:   os.Getenv("CREATE_RATING_REPLY_SUBJECT"),
		DeleteRatingCommandSubject: os.Getenv("DELETE_RATING_COMMAND_SUBJECT"),
		DeleteRatingReplySubject:   os.Getenv("DELETE_RATING_REPLY_SUBJECT"),
		NotificationSubject:        os.Getenv("NOTIFICATION_SUBJECT"),
	}
}
