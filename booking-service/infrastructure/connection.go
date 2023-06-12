package infrastructure

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

//func GetAuthClient(cfg *config.Config) *auth.AuthServiceClient {
//	conn := createConnection((*cfg).AuthServiceAddress)
//	client := auth.NewAuthServiceClient(conn)
//	return &client
//}

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
