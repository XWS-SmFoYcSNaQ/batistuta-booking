package infrastructure

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateConnection(address string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
