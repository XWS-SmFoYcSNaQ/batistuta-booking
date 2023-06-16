package client

import (
	"booking_service/config"
	"context"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/accommodation"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"regexp"
)

type AccommodationServiceClient struct {
	Client accommodation.AccommodationServiceClient
}

func InitAccommodationServiceClient(url string) (AccommodationServiceClient, error) {
	cc, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return AccommodationServiceClient{Client: nil}, err
	}

	c := AccommodationServiceClient{
		Client: accommodation.NewAccommodationServiceClient(cc),
	}

	return c, nil
}

func GetServer() *http.Server {
	cfg := config.LoadConfig()
	gwmux := runtime.NewServeMux()

	RegisterAccommodationClient(gwmux, &cfg)
	log.Println("ADDRESS" + cfg.AccommodationServiceAddress)

	return &http.Server{
		Addr:    cfg.AccommodationServiceAddress,
		Handler: cors(gwmux),
	}
}

func allowedOrigin(origin string) bool {
	if viper.GetString("cors") == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
		return true
	}
	return false
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func RegisterAccommodationClient(Mux *runtime.ServeMux, Cfg *config.Config) accommodation.AccommodationServiceClient {
	conn, err := createConnection(Cfg.AccommodationServiceAddress)
	handleError(err, "Failed to create connection with address: "+Cfg.AccommodationServiceAddress)
	client := accommodation.NewAccommodationServiceClient(conn)
	err = accommodation.RegisterAccommodationServiceHandlerClient(
		context.Background(),
		Mux,
		client,
	)
	handleError(err, "Failed to register Accommodation Microservice")
	return client
}

func createConnection(address string) (*grpc.ClientConn, error) {
	conn, err := grpc.DialContext(
		context.Background(),
		address,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (c *AccommodationServiceClient) GetAllAccommodationsForHost(ctx context.Context) (*accommodation.AM_GetAllAccommodations_Response, error) {
	req := &accommodation.AM_GetAllAccommodations_Request{}
	return c.Client.GetMyAccommodations(ctx, req)
}

func handleError(e error, m string) {
	if e != nil {
		log.Fatalln(m, e)
	}
}
