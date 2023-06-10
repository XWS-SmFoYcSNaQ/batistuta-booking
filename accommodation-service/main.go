package main

import (
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/config"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/controller"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/handlers"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/infrastructure"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/infrastructure/database"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/services"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/accommodation"
	commonServices "github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/services"
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

	authClient := infrastructure.GetAuthClient(&cfg)

	accommodationHandler := handlers.AccommodationHandler{
		AccommodationController: &controller.AccommodationController{
			AccommodationService: &services.AccommodationService{DB: db},
			PeriodService:        &services.PeriodService{DB: db},
			DiscountService:      &services.DiscountService{DB: db},
			AuthService: &commonServices.AuthService{
				AuthClient: authClient,
			},
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
