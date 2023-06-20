package infrastructure

import (
	"context"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/config"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func GetAuthClient(cfg *config.Config) *auth.AuthServiceClient {
	conn := CreateConnection((*cfg).AuthServiceAddress)
	client := auth.NewAuthServiceClient(conn)
	return &client
}

func CreateConnection(address string) *grpc.ClientConn {
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
