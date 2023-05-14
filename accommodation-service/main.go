package main

import (
	"accommodation_service/config"
	"accommodation_service/controller"
	"accommodation_service/database"
	"accommodation_service/handlers"
	"accommodation_service/proto/accommodation"
	"accommodation_service/services"
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

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	accommodationHandler := handlers.AccommodationHandler{
		AccommodationController: &controller.AccommodationController{
			AccommodationService: &services.AccommodationService{DB: db},
		},
		PeriodController: &controller.PeriodController{
			PeriodService: &services.PeriodService{DB: db},
		},
		DiscountController: &controller.DiscountController{
			DiscountService: &services.DiscountService{DB: db},
		},
	}
	accommodation.RegisterAccommodationServiceServer(grpcServer, accommodationHandler)

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
