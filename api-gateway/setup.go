package main

import (
	"api_gateway/client"
	"api_gateway/config"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"regexp"
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
	log.Println("ADDRESS" + cfg.Address)

	return &http.Server{
		Addr:    cfg.Address,
		Handler: cors(gwmux),
	}
}

func allowedOrigin(origin string) bool {
	if viper.GetString("cors") == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
		return true
	}
	return false
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
