package main

import (
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/config"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/startup"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
