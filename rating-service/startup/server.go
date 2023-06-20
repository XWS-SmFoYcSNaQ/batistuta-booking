package startup

import (
	"fmt"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/messaging"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/messaging/nats"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/auth"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/rating"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/services"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/application"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/config"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/domain"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/infrastructure"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/infrastructure/database"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/infrastructure/database/mock"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	CreateQueueGroup = "rating_service_create"
	DeleteQueueGroup = "rating_service_delete"
)

func (server *Server) Start() {
	postgresClient := server.initPostgresClient()
	ratingRepo := server.initRatingRepository(postgresClient)
	authService := services.NewAuthService(server.GetAuthClient())

	notificationPublisher := server.initPublisher(server.config.NotificationSubject)

	createCommandPublisher := server.initPublisher(server.config.CreateRatingCommandSubject)
	createReplySubscriber := server.initSubscriber(server.config.CreateRatingReplySubject, CreateQueueGroup)
	createRatingOrchestrator := server.initCreateRatingOrchestrator(&createCommandPublisher, &createReplySubscriber)

	createCommandSubscriber := server.initSubscriber(server.config.CreateRatingCommandSubject, CreateQueueGroup)
	createReplyPublisher := server.initPublisher(server.config.CreateRatingReplySubject)

	deleteCommandPublisher := server.initPublisher(server.config.DeleteRatingCommandSubject)
	deleteReplySubscriber := server.initSubscriber(server.config.DeleteRatingReplySubject, DeleteQueueGroup)
	deleteRatingOrchestrator := server.initDeleteRatingOrchestrator(&deleteCommandPublisher, &deleteReplySubscriber)

	deleteCommandSubscriber := server.initSubscriber(server.config.DeleteRatingCommandSubject, DeleteQueueGroup)
	deleteReplyPublisher := server.initPublisher(server.config.DeleteRatingReplySubject)

	productService := server.initRatingService(ratingRepo, createRatingOrchestrator, deleteRatingOrchestrator)
	server.initCreateRatingHandler(productService, &createReplyPublisher, &createCommandSubscriber, &notificationPublisher)
	server.initDeleteRatingHandler(productService, &deleteReplyPublisher, &deleteCommandSubscriber, &notificationPublisher)

	productHandler := server.initRatingHandler(productService, authService)
	server.startGrpcServer(productHandler)
}

func (server *Server) initPostgresClient() *gorm.DB {
	client, err := database.GetClient(
		server.config.RatingDBHost, server.config.RatingDBUser,
		server.config.RatingDBPass, server.config.RatingDBName,
		server.config.RatingDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initRatingRepository(client *gorm.DB) *domain.RatingRepository {
	store, err := database.NewRatingPostgresRepository(client)
	if err != nil {
		log.Fatal(err)
	}
	store.DeleteAll()
	for _, r := range mock.Ratings {
		err := store.Insert(r)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &store
}

func (server *Server) initRatingService(
	repository *domain.RatingRepository,
	createOrchestrator *domain.CreateRatingOrchestrator,
	deleteOrchestrator *domain.DeleteRatingOrchestrator) *domain.RatingService {
	return domain.NewRatingService(repository, createOrchestrator, deleteOrchestrator)
}

func (server *Server) initPublisher(subject string) messaging.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) messaging.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initCreateRatingHandler(service *domain.RatingService, publisher *messaging.Publisher, subscriber *messaging.Subscriber, notificationPublisher *messaging.Publisher) {
	_, err := application.NewCreateRatingCommandHandler(service, publisher, subscriber, notificationPublisher)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initDeleteRatingHandler(service *domain.RatingService, publisher *messaging.Publisher, subscriber *messaging.Subscriber, notificationPublisher *messaging.Publisher) {
	_, err := application.NewDeleteRatingCommandHandler(service, publisher, subscriber, notificationPublisher)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initCreateRatingOrchestrator(publisher *messaging.Publisher, subscriber *messaging.Subscriber) *domain.CreateRatingOrchestrator {
	orchestrator, err := domain.NewCreateRatingOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initDeleteRatingOrchestrator(publisher *messaging.Publisher, subscriber *messaging.Subscriber) *domain.DeleteRatingOrchestrator {
	orchestrator, err := domain.NewDeleteRatingOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initRatingHandler(service *domain.RatingService, authService *services.AuthService) *application.RatingHandler {
	return application.NewRatingHandler(server.config, service, authService)
}

func (server *Server) startGrpcServer(ratingHandler *application.RatingHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	rating.RegisterRatingServiceServer(grpcServer, ratingHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (server *Server) GetAuthClient() *auth.AuthServiceClient {
	conn := infrastructure.CreateConnection((*server).config.AuthServiceAddress)
	client := auth.NewAuthServiceClient(conn)
	return &client
}
