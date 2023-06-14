package main

import (
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/config"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/controller"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/handlers"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/infrastructure"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/infrastructure/database"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/services"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/messaging"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/messaging/nats"
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

const (
	CreateQueueGroup = "accommodation_service_create"
	DeleteQueueGroup = "accommodation_service_delete"
)

func initSubscriber(config config.Config, subject, queueGroup string) messaging.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		config.NatsHost, config.NatsPort,
		config.NatsUser, config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func initPublisher(config config.Config, subject string) messaging.Publisher {
	publisher, err := nats.NewNATSPublisher(
		config.NatsHost, config.NatsPort,
		config.NatsUser, config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func initCreateRatingHandler(service *services.RatingService, publisher *messaging.Publisher, subscriber *messaging.Subscriber) {
	_, err := handlers.NewCreateRatingCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func initDeleteRatingHandler(service *services.RatingService, publisher *messaging.Publisher, subscriber *messaging.Subscriber) {
	_, err := handlers.NewDeleteRatingCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

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

	createCommandSubscriber := initSubscriber(cfg, cfg.CreateRatingCommandSubject, CreateQueueGroup)
	createReplyPublisher := initPublisher(cfg, cfg.CreateRatingReplySubject)
	deleteCommandSubscriber := initSubscriber(cfg, cfg.DeleteRatingCommandSubject, DeleteQueueGroup)
	deleteReplyPublisher := initPublisher(cfg, cfg.DeleteRatingReplySubject)
	initCreateRatingHandler(&services.RatingService{DB: db}, &createReplyPublisher, &createCommandSubscriber)
	initDeleteRatingHandler(&services.RatingService{DB: db}, &deleteReplyPublisher, &deleteCommandSubscriber)

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
			AuthService: &commonServices.AuthService{
				AuthClient: authClient,
			},
		},
		DiscountController: &controller.DiscountController{
			DiscountService: &services.DiscountService{DB: db},
			AuthService: &commonServices.AuthService{
				AuthClient: authClient,
			},
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
