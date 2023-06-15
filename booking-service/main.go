package main

import (
	"booking_service/config"
	"booking_service/controller"
	"booking_service/handlers"
	"booking_service/infrastructure/database"
	"booking_service/proto/booking"
	"booking_service/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.LoadConfig()

	db := database.Connect(cfg)
	defer db.Close()

	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		log.Fatalln(err)
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(listener)

	// Bootstrap gRPC server.
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	//accommodationClient := infrastructure.GetAccommodationClient(&cfg)
	//authClient := infrastructure.GetAuthClient(&cfg)
	// Bootstrap gRPC service server and respond to request.
	bookingHandler := handlers.BookingHandler{
		ReservationController: &controller.ReservationController{
			BookingService: &services.BookingRequestsService{DB: db},
		},
	}
	booking.RegisterBookingServiceServer(grpcServer, bookingHandler)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	grpcServer.Stop()
}
