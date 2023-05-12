package main

import (
	"api_gateway/client"
	"api_gateway/config"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func registerServices(gwmux *runtime.ServeMux, cfg *config.Config) {
	client.RegisterAccommodationClient(gwmux, cfg)
	client.RegisterBookingClient(gwmux, cfg)
	client.RegisterUserClient(gwmux, cfg)
	client.RegisterAuthClient(gwmux, cfg)
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
