package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"recommendation-service/config"
	"recommendation-service/controller"
	"recommendation-service/handlers"
	"recommendation-service/infrastructure/database"
	"recommendation-service/proto/recommendation"
	"recommendation-service/services"
	"syscall"
)

func main() {
	cfg := config.LoadConfig()
	//Initialize the logger we are going to use, with prefix and datetime for every log
	storeLogger := log.New(os.Stdout, "[recommendation-service] ", log.LstdFlags)

	// NoSQL: Initialize Movie Repository store
	db, err := database.Create(cfg, storeLogger)
	if err != nil {
		storeLogger.Fatal(err)
	}
	defer db.CloseDriverConnection(context.Background())
	db.CheckConnection()
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)
	storeLogger.Println("Successfully connected to neo4j database!")

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

	recommendationHandler := handlers.RecommendationHandler{
		RecommendationController: &controller.RecommendationController{
			RecommendationService: &services.RecommendationService{RecommendationRepo: db},
		},
	}

	recommendation.RegisterRecommendationServiceServer(grpcServer, recommendationHandler)

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
