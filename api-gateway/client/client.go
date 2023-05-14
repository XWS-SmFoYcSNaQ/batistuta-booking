package client

import (
	"api_gateway/config"
	"api_gateway/proto/accommodation"
	"api_gateway/proto/auth"
	"api_gateway/proto/booking"
	"api_gateway/proto/user"
	"context"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterAccommodationClient(Mux *runtime.ServeMux, Cfg *config.Config) {
	conn := createConnection(Cfg.AccommodationServiceAddress)
	client := accommodation.NewAccommodationServiceClient(conn)
	err := accommodation.RegisterAccommodationServiceHandlerClient(
		context.Background(),
		Mux,
		client,
	)
	handleError(err, "Failed to register Accommodation Microservice")
}

func RegisterBookingClient(Mux *runtime.ServeMux, Cfg *config.Config) {
	conn := createConnection(Cfg.BookingServiceAddress)
	client := booking.NewBookingServiceClient(conn)
	err := booking.RegisterBookingServiceHandlerClient(
		context.Background(),
		Mux,
		client,
	)
	handleError(err, "Failed to register Booking Microservice")
}

func RegisterUserClient(Mux *runtime.ServeMux, Cfg *config.Config) {
	conn := createConnection(Cfg.UserServiceAddress)
	client := user.NewUserServiceClient(conn)
	err := user.RegisterUserServiceHandlerClient(
		context.Background(),
		Mux,
		client,
	)
	handleError(err, "Failed to register User Microservice")
}

func RegisterAuthClient(Mux *runtime.ServeMux, Cfg *config.Config) {
	conn := createConnection(Cfg.AuthServiceAddress)
	client := auth.NewAuthServiceClient(conn)
	err := auth.RegisterAuthServiceHandlerClient(
		context.Background(),
		Mux,
		client,
	)
	handleError(err, "Failed to register Auth Microservice")
}

// private

func createConnection(address string) *grpc.ClientConn {
	conn, err := grpc.DialContext(
		context.Background(),
		address,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	handleError(err, "Failed to create connection with address: "+address)
	return conn
}

func handleError(e error, m string) {
	if e != nil {
		log.Fatalln(m, e)
	}
}
