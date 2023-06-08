package startup

import (
	"fmt"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/rating"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/application"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/domain"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/infrastructure/database"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/rating_service/infrastructure/database/mock"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
)

type Server struct {
	config *Config
}

func NewServer(config *Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "rating_service"
)

func (server *Server) Start() {
	postgresClient := server.initPostgresClient()
	ratingRepo := server.initRatingRepository(postgresClient)

	productService := server.initRatingService(ratingRepo)

	//commandSubscriber := server.initSubscriber(server.config.CreateOrderCommandSubject, QueueGroup)
	//replyPublisher := server.initPublisher(server.config.CreateOrderReplySubject)
	//server.initCreateOrderHandler(productService, replyPublisher, commandSubscriber)

	productHandler := server.initRatingHandler(productService)

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

func (server *Server) initRatingRepository(client *gorm.DB) domain.RatingRepository {
	store, err := database.NewRatingPostgresRepository(client)
	if err != nil {
		log.Fatal(err)
	}
	store.DeleteAll()
	for _, Product := range mock.Ratings {
		err := store.Insert(Product)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initRatingService(repository domain.RatingRepository) *domain.RatingService {
	return domain.NewRatingService(&repository)
}

func (server *Server) initRatingHandler(service *domain.RatingService) *application.RatingHandler {
	return application.NewRatingHandler(service)
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
