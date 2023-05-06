package main

import (
	"api_gateway/client"
	"api_gateway/config"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

func registerServices(gwmux *runtime.ServeMux, cfg *config.Config) {
	client.RegisterAccommodationClient(gwmux, cfg)
	client.RegisterBookingClient(gwmux, cfg)
}

func GetServer() *http.Server {
	cfg := config.LoadConfig()
	gwmux := runtime.NewServeMux()

	registerServices(gwmux, &cfg)

	return &http.Server{
		Addr:    cfg.Address,
		Handler: gwmux,
	}
}
