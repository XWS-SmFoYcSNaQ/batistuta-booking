package controller

import (
	"booking_service/infrastructure"
	"booking_service/model"
	"booking_service/proto/booking"
	"booking_service/services"
	"context"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/accommodation"
	commonServices "github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/services"
	"google.golang.org/grpc/metadata"
	"os"

	//commonServices "github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ReservationController struct {
	BookingService *services.BookingRequestsService
	AuthService    *commonServices.AuthService
}

func (c ReservationController) GetAll(ctx context.Context, request *booking.AM_GetAllBookingRequests_Request) (*booking.AM_GetAllBookingRequests_Response, error) {
	bookingRequests, err := c.BookingService.GetAll()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	var res []*booking.BookingRequestsDTO
	for _, d := range bookingRequests {
		a := booking.BookingRequestsDTO{
			Id:              d.ID.String(),
			AccommodationId: d.AccommodationId,
			StartDate:       d.StartDate,
			EndDate:         d.EndDate,
			NumberOfGuests:  int32(d.NumberOfGuests),
			UserId:          d.UserId,
		}
		res = append(res, &a)
	}

	return &booking.AM_GetAllBookingRequests_Response{Data: res}, nil
}

func (c ReservationController) GetAllByUserId(ctx context.Context, request *booking.AM_GetAllBookingRequestsByUserId_Request) (*booking.AM_GetAllBookingRequests_Response, error) {
	bookingRequests, err := c.BookingService.GetAllByUserId(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	var res []*booking.BookingRequestsDTO
	for _, d := range bookingRequests {
		a := booking.BookingRequestsDTO{
			Id:              d.ID.String(),
			AccommodationId: d.AccommodationId,
			StartDate:       d.StartDate,
			EndDate:         d.EndDate,
			NumberOfGuests:  int32(d.NumberOfGuests),
			UserId:          d.UserId,
		}
		res = append(res, &a)
	}

	return &booking.AM_GetAllBookingRequests_Response{Data: res}, nil
}

func (c ReservationController) MakeBookingRequest(ctx context.Context, request *booking.AM_BookingRequest_Request) (*booking.AM_CreateBookingRequest_Response, error) {
	id, err := c.BookingService.MakeBookingRequest(&model.BookingRequest{
		AccommodationId: request.AccommodationId,
		StartDate:       request.StartDate,
		EndDate:         request.EndDate,
		UserId:          request.UserId,
		NumberOfGuests:  int(request.NumberOfGuests),
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &booking.AM_CreateBookingRequest_Response{Id: id.String()}, nil
}

func (c ReservationController) DeleteBookingRequest(ctx context.Context, request *booking.AM_DeleteBookingRequest_Request) (*booking.AM_DeleteBookingRequest_Response, error) {
	err := c.BookingService.DeleteBookingRequest(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &booking.AM_DeleteBookingRequest_Response{}, nil
}

func (c ReservationController) ConfirmReservationRequest(ctx context.Context, request *booking.ReservationConfirm_Request) (*booking.EmptyMessage, error) {
	err := c.BookingService.ConfirmReservation(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &booking.EmptyMessage{}, nil
}

func (c ReservationController) GetAllReservationsForGuest(ctx context.Context, request *booking.AllReservationsForGuest_Request) (*booking.AllReservationsForGuest_Response, error) {
	bookingRequests, err := c.BookingService.GetAllReservationsForUser(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	var res []*booking.BookingRequestsDTO
	for _, d := range bookingRequests {
		a := booking.BookingRequestsDTO{
			Id:              d.ID.String(),
			AccommodationId: d.AccommodationId,
			StartDate:       d.StartDate,
			EndDate:         d.EndDate,
			NumberOfGuests:  int32(d.NumberOfGuests),
			UserId:          d.UserId,
		}
		res = append(res, &a)
	}

	return &booking.AllReservationsForGuest_Response{Data: res}, nil
}

func (c ReservationController) DeleteReservation(ctx context.Context, request *booking.DeleteReservation_Request) (*booking.EmptyMessage, error) {
	err := c.BookingService.DeleteReservation(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &booking.EmptyMessage{}, nil
}

func (c ReservationController) GetReservationsForHost(ctx context.Context, request *booking.EmptyMessage) (*booking.ReservationsForHost_Response, error) {
	// Get the authorization header from the incoming context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing authorization header")
	}
	authHeaders := md.Get("Authorization")
	if len(authHeaders) == 0 {
		return nil, status.Error(codes.Unauthenticated, "missing authorization header")
	}
	authHeader := authHeaders[0]

	// Create a gRPC connection to the API Gateway server
	conn, err := infrastructure.CreateConnection(os.Getenv("ACCOMMODATION_SERVICE_ADDRESS"))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// Create the gRPC client by specifying the registered client from the API Gateway
	client := accommodation.NewAccommodationServiceClient(conn)

	// Create a new context with the authorization header
	authCtx := metadata.AppendToOutgoingContext(ctx, "Authorization", authHeader)

	// Make the gRPC call with the updated context
	response, err := client.GetMyAccommodations(authCtx, &accommodation.AM_GetAllAccommodations_Request{})
	if err != nil {
		return nil, err
	}

	// Extract accommodationIds from the response
	var accommodationIDs []string
	for _, accommodationDTO := range response.Data {
		accommodationIDs = append(accommodationIDs, accommodationDTO.Id)
	}

	// Use accommodationIDs to get reservations
	reservations, err := c.BookingService.GetReservationsForAccommodationIDs(accommodationIDs)
	if err != nil {
		return nil, err
	}

	var res []*booking.BookingRequestsDTO

	// Populate the response with the reservation data
	for _, reservation := range reservations {
		bookingDTO := booking.BookingRequestsDTO{
			Id:              reservation.ID.String(),
			AccommodationId: reservation.AccommodationId,
			StartDate:       reservation.StartDate,
			EndDate:         reservation.EndDate,
			UserId:          reservation.UserId,
			NumberOfGuests:  int32(reservation.NumberOfGuests),
		}
		res = append(res, &bookingDTO)
	}

	return &booking.ReservationsForHost_Response{Data: res}, nil
}

func (c ReservationController) GetReservationRequestsForHost(ctx context.Context, request *booking.EmptyMessage) (*booking.ReservationsForHost_Response, error) {
	// Get the authorization header from the incoming context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing authorization header")
	}
	authHeaders := md.Get("Authorization")
	if len(authHeaders) == 0 {
		return nil, status.Error(codes.Unauthenticated, "missing authorization header")
	}
	authHeader := authHeaders[0]

	// Create a gRPC connection to the API Gateway server
	conn, err := infrastructure.CreateConnection(os.Getenv("ACCOMMODATION_SERVICE_ADDRESS"))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// Create the gRPC client by specifying the registered client from the API Gateway
	client := accommodation.NewAccommodationServiceClient(conn)

	// Create a new context with the authorization header
	authCtx := metadata.AppendToOutgoingContext(ctx, "Authorization", authHeader)

	// Make the gRPC call with the updated context
	response, err := client.GetMyAccommodations(authCtx, &accommodation.AM_GetAllAccommodations_Request{})
	if err != nil {
		return nil, err
	}

	// Extract accommodationIds from the response
	var accommodationIDs []string
	for _, accommodationDTO := range response.Data {
		accommodationIDs = append(accommodationIDs, accommodationDTO.Id)
	}

	// Use accommodationIDs to get reservations
	reservations, err := c.BookingService.GetReservationRequestsForAccommodationIDs(accommodationIDs)
	if err != nil {
		return nil, err
	}

	var res []*booking.BookingRequestsDTO

	// Populate the response with the reservation data
	for _, reservation := range reservations {
		bookingDTO := booking.BookingRequestsDTO{
			Id:                           reservation.ID.String(),
			AccommodationId:              reservation.AccommodationId,
			StartDate:                    reservation.StartDate,
			EndDate:                      reservation.EndDate,
			UserId:                       reservation.UserId,
			NumberOfGuests:               int32(reservation.NumberOfGuests),
			NumberOfCanceledReservations: c.BookingService.GetNumberOfCanceledReservationsForGuest(reservation.UserId),
		}
		res = append(res, &bookingDTO)
	}

	return &booking.ReservationsForHost_Response{Data: res}, nil
}

func (c ReservationController) HostStandOutCheck(ctx context.Context, request *booking.EmptyMessage) (*booking.StandOutHost_Response, error) {
	// Get the authorization header from the incoming context
	// Get the authorization header from the incoming context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing authorization header")
	}
	authHeaders := md.Get("Authorization")
	if len(authHeaders) == 0 {
		return nil, status.Error(codes.Unauthenticated, "missing authorization header")
	}
	authHeader := authHeaders[0]

	// Create a gRPC connection to the API Gateway server
	conn, err := infrastructure.CreateConnection(os.Getenv("ACCOMMODATION_SERVICE_ADDRESS"))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// Create the gRPC client by specifying the registered client from the API Gateway
	client := accommodation.NewAccommodationServiceClient(conn)

	// Create a new context with the authorization header
	authCtx := metadata.AppendToOutgoingContext(ctx, "Authorization", authHeader)

	// Make the gRPC call with the updated context
	response, err := client.GetMyAccommodations(authCtx, &accommodation.AM_GetAllAccommodations_Request{})
	if err != nil {
		return nil, err
	}

	// Extract accommodationIds from the response
	var accommodationIDs []string
	for _, accommodationDTO := range response.Data {
		accommodationIDs = append(accommodationIDs, accommodationDTO.Id)
	}

	if len(accommodationIDs) == 0 {
		return &booking.StandOutHost_Response{Flag: false, Message: "Not presented accommodations"}, nil
	}

	resp, err := c.BookingService.IsTheCancellationRateLessThanFive(accommodationIDs)
	if err != nil {
		return nil, err
	}
	finalResponse := &booking.StandOutHost_Response{Flag: true, Message: ""}
	if resp == false {
		finalResponse.Flag = false
		finalResponse.Message = "Has a cancellation rate greater than 5%."
	}
	resp, err = c.BookingService.HasAtLeastFivePastReservations(accommodationIDs)
	if err != nil {
		return nil, err
	}
	if resp == false {
		finalResponse.Flag = false
		finalResponse.Message += "Has had less than 5 reservations for accommodations in the past."
	}
	resp, err = c.BookingService.IsTotalReservationDurationGreaterThanFiftyDays(accommodationIDs)
	if err != nil {
		return nil, err
	}
	if resp == false {
		finalResponse.Flag = false
		finalResponse.Message += "The total duration of all reservations is less than 50 days."
	}
	return &booking.StandOutHost_Response{Flag: finalResponse.Flag, Message: finalResponse.Message}, nil
}
