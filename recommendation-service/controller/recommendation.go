package controller

import (
	"context"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/accommodation"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/booking"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/rating"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"recommendation-service/infrastructure"
	"recommendation-service/proto/recommendation"
	"recommendation-service/services"
	"strconv"
)

type RecommendationController struct {
	RecommendationService *services.RecommendationService
}

//func (c RecommendationController) Test(ctx context.Context, request *recommendation.Test_Request) (*recommendation.Test_Response, error) {
//	return c.RecommendationService.Test()
//}

func (c RecommendationController) GetRecommendedAccommodations(ctx context.Context, request *recommendation.RecommendedAccommodations_Request) (*recommendation.RecommendedAccommodations_Response, error) {
	// Get the authorization header from the incoming context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing authorization header")
	}
	authHeaders := md.Get("Authorization")
	if len(authHeaders) == 0 {
		return nil, status.Error(codes.Unauthenticated, "missing authorization header")
	}
	//authHeader := authHeaders[0]

	// Create a gRPC connection to the API Gateway server
	accommodationConn, err := infrastructure.CreateConnection(os.Getenv("ACCOMMODATION_SERVICE_ADDRESS"))
	if err != nil {
		return nil, err
	}
	defer accommodationConn.Close()

	// Create the gRPC client by specifying the registered client from the API Gateway
	accommodationClient := accommodation.NewAccommodationServiceClient(accommodationConn)

	//// Create a new context with the authorization header
	//authCtx := metadata.AppendToOutgoingContext(ctx, "Authorization", authHeader)

	// Make the gRPC call with the updated context
	accommodationsResponse, err := accommodationClient.GetAllAccommodations(ctx, &accommodation.AM_GetAllAccommodations_Request{})
	if err != nil {
		return nil, err
	}

	userConn, err := infrastructure.CreateConnection(os.Getenv("USER_SERVICE_ADDRESS"))
	if err != nil {
		return nil, err
	}
	defer userConn.Close()

	// Create the gRPC client by specifying the registered client from the API Gateway
	userClient := user.NewUserServiceClient(userConn)

	//// Create a new context with the authorization header
	//authCtx := metadata.AppendToOutgoingContext(ctx, "Authorization", authHeader)

	// Make the gRPC call with the updated context
	usersResponse, err := userClient.GetAllUsers(ctx, &user.Empty_Message{})
	if err != nil {
		return nil, err
	}
	// Create a gRPC connection to the API Gateway server

	ratingConn, err := infrastructure.CreateConnection(os.Getenv("RATING_SERVICE_ADDRESS"))
	if err != nil {
		return nil, err
	}
	defer userConn.Close()

	ratingClient := rating.NewRatingServiceClient(ratingConn)
	ratingResponse, err := ratingClient.GetAllRatings(ctx, &rating.Empty{})

	bookingConn, err := infrastructure.CreateConnection(os.Getenv("BOOKING_SERVICE_ADDRESS"))
	if err != nil {
		return nil, err
	}
	defer bookingConn.Close()

	bookingClient := booking.NewBookingServiceClient(bookingConn)

	var allReservations []*booking.BookingRequestsDTO

	for _, userIter := range usersResponse.Users {
		request := &booking.AllReservationsForGuest_Request{
			Id: userIter.Id,
		}

		response, err := bookingClient.GetAllReservationsForGuest(ctx, request)
		if err != nil {
			return nil, err
		}

		allReservations = append(allReservations, response.Data...)
	}

	log.Println("Broj svih rezervacija je " + strconv.Itoa(len(allReservations)))
	c.RecommendationService.SetDataForDb(accommodationsResponse.Data, usersResponse.Users, ratingResponse.Data, allReservations)

	recommendedAccommodations, err := c.RecommendationService.GetRecommendedAccommodations(request.Id)

	return &recommendation.RecommendedAccommodations_Response{Data: recommendedAccommodations}, nil
}
