package infrastructure

import (
	"booking_service/config"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

//func GetAccommodationClient(cfg *config.Config) *accommodation.AccommodationServiceClient {
//	conn := CreateConnection((*cfg).AccommodationServiceAddress)
//	client := accommodation.NewAccommodationServiceClient(conn)
//	return &client
//}

func GetAuthClient(cfg *config.Config) *auth.AuthServiceClient {
	conn, err := CreateConnection((*cfg).AuthServiceAddress)
	handleError(err, "Couldn't create auth client")
	client := auth.NewAuthServiceClient(conn)
	return &client
}

func CreateConnection(address string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func handleError(e error, m string) {
	if e != nil {
		log.Fatalln(m, e)
	}
}
