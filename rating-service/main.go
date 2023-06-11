package main

import (
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/startup"
)

func main() {
	config := startup.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
