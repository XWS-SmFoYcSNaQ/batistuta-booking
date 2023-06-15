package infrastructure

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

//func GetAccommodationClient(cfg *config.Config) *accommodation.AccommodationServiceClient {
//	conn := CreateConnection((*cfg).AccommodationServiceAddress)
//	client := accommodation.NewAccommodationServiceClient(conn)
//	return &client
//}

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
